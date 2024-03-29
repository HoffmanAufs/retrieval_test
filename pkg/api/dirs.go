// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gauss-project/aurorafs/pkg/boson"
	"github.com/gauss-project/aurorafs/pkg/collection/entry"
	"github.com/gauss-project/aurorafs/pkg/file"
	"github.com/gauss-project/aurorafs/pkg/file/loadsave"
	"github.com/gauss-project/aurorafs/pkg/jsonhttp"
	"github.com/gauss-project/aurorafs/pkg/logging"
	"github.com/gauss-project/aurorafs/pkg/manifest"
	"github.com/gauss-project/aurorafs/pkg/tracing"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeTar    = "application/x-tar"
)

const (
	manifestRootPath                      = "/"
	manifestWebsiteIndexDocumentSuffixKey = "website-index-document"
	manifestWebsiteErrorDocumentPathKey   = "website-error-document"
)

// dirUploadHandler uploads a directory supplied as a tar in an HTTP request
func (s *server) dirUploadHandler(w http.ResponseWriter, r *http.Request) {
	logger := tracing.NewLoggerWithTraceID(r.Context(), s.logger)
	err := validateRequest(r)
	if err != nil {
		logger.Errorf("dir upload, validate request")
		logger.Debugf("dir upload, validate request err: %v", err)
		jsonhttp.BadRequest(w, "could not validate request")
		return
	}



	// Add the tag to the context
	ctx := r.Context()
	p := requestPipelineFn(s.storer, r)
	encrypt := requestEncrypt(r)
	l := loadsave.New(s.storer, requestModePut(r), encrypt)
	reference, err := storeDir(ctx, encrypt, r.Body, s.logger, p, l, r.Header.Get(BosonIndexDocumentHeader), r.Header.Get(BosonErrorDocumentHeader))
	if err != nil {
		logger.Debugf("dir upload: store dir err: %v", err)
		logger.Errorf("dir upload: store dir")
		jsonhttp.InternalServerError(w, "could not store dir")
		return
	}

	a, err := s.traversal.GetTrieData(ctx, reference)
	if err != nil {
		logger.Errorf("dir upload: get trie data err: %v", err)
		jsonhttp.InternalServerError(w, "could not get trie data")
		return
	}
	dataChunks ,_ := s.traversal.CheckTrieData(ctx, reference, a)
	if err != nil {
		logger.Errorf("dir upload: check trie data err: %v", err)
		jsonhttp.InternalServerError(w, "check trie data error")
		return
	}
	for _,li := range dataChunks {
		for _,b := range li {
			s.chunkInfo.OnChunkTransferred(boson.NewAddress(b), reference, s.overlay)
		}
	}

	jsonhttp.OK(w, fileUploadResponse{
		Reference: reference,
	})
}

// validateRequest validates an HTTP request for a directory to be uploaded
func validateRequest(r *http.Request) error {
	if r.Body == http.NoBody {
		return errors.New("request has no body")
	}
	contentType := r.Header.Get(contentTypeHeader)
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != contentTypeTar {
		return errors.New("content-type not set to tar")
	}
	return nil
}

// storeDir stores all files recursively contained in the directory given as a tar
// it returns the hash for the uploaded manifest corresponding to the uploaded dir
func storeDir(ctx context.Context, encrypt bool, reader io.ReadCloser, log logging.Logger, p pipelineFunc, ls file.LoadSaver, indexFilename string, errorFilename string) (boson.Address, error) {
	logger := tracing.NewLoggerWithTraceID(ctx, log)

	dirManifest, err := manifest.NewDefaultManifest(ls, encrypt)
	if err != nil {
		return boson.ZeroAddress, err
	}

	if indexFilename != "" && strings.ContainsRune(indexFilename, '/') {
		return boson.ZeroAddress, fmt.Errorf("index document suffix must not include slash character")
	}

	// set up HTTP body reader
	tarReader := tar.NewReader(reader)
	defer reader.Close()

	filesAdded := 0

	// iterate through the files in the supplied tar
	for {
		fileHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return boson.ZeroAddress, fmt.Errorf("read tar stream: %w", err)
		}

		filePath := filepath.Clean(fileHeader.Name)

		if filePath == "." {
			logger.Warning("skipping file upload empty path")
			continue
		}

		if runtime.GOOS == "windows" {
			// always use Unix path separator
			filePath = filepath.ToSlash(filePath)
		}

		// only store regular files
		if !fileHeader.FileInfo().Mode().IsRegular() {
			logger.Warningf("skipping file upload for %s as it is not a regular file", filePath)
			continue
		}

		fileName := fileHeader.FileInfo().Name()
		contentType := mime.TypeByExtension(filepath.Ext(fileHeader.Name))

		// upload file
		fileInfo := &fileUploadInfo{
			name:        fileName,
			size:        fileHeader.FileInfo().Size(),
			contentType: contentType,
			reader:      tarReader,
		}

		fileReference, err := storeFile(ctx, fileInfo, p, encrypt)
		if err != nil {
			return boson.ZeroAddress, fmt.Errorf("store dir file: %w", err)
		}
		logger.Tracef("uploaded dir file %v with reference %v", filePath, fileReference)

		// add file entry to dir manifest
		err = dirManifest.Add(ctx, filePath, manifest.NewEntry(fileReference, nil))
		if err != nil {
			return boson.ZeroAddress, fmt.Errorf("add to manifest: %w", err)
		}

		filesAdded++
	}

	// check if files were uploaded through the manifest
	if filesAdded == 0 {
		return boson.ZeroAddress, fmt.Errorf("no files in tar")
	}

	// store website information
	if indexFilename != "" || errorFilename != "" {
		metadata := map[string]string{}
		if indexFilename != "" {
			metadata[manifestWebsiteIndexDocumentSuffixKey] = indexFilename
		}
		if errorFilename != "" {
			metadata[manifestWebsiteErrorDocumentPathKey] = errorFilename
		}
		rootManifestEntry := manifest.NewEntry(boson.ZeroAddress, metadata)
		err = dirManifest.Add(ctx, manifestRootPath, rootManifestEntry)
		if err != nil {
			return boson.ZeroAddress, fmt.Errorf("add to manifest: %w", err)
		}
	}

	storeSizeFn := []manifest.StoreSizeFunc{}

	// save manifest
	manifestBytesReference, err := dirManifest.Store(ctx, storeSizeFn...)
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("store manifest: %w", err)
	}

	// store the manifest metadata and get its reference
	m := entry.NewMetadata(manifestBytesReference.String())
	m.MimeType = dirManifest.Type()
	metadataBytes, err := json.Marshal(m)
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("metadata marshal: %w", err)
	}


	mr, err := p(ctx, bytes.NewReader(metadataBytes), int64(len(metadataBytes)))
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("split metadata: %w", err)
	}

	// now join both references (fr, mr) to create an entry and store it
	e := entry.New(manifestBytesReference, mr)
	fileEntryBytes, err := e.MarshalBinary()
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("entry marshal: %w", err)
	}

	manifestFileReference, err := p(ctx, bytes.NewReader(fileEntryBytes), int64(len(fileEntryBytes)))
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("split entry: %w", err)
	}

	return manifestFileReference, nil
}

// storeFile uploads the given file and returns its reference
// this function was extracted from `fileUploadHandler` and should eventually replace its current code
func storeFile(ctx context.Context, fileInfo *fileUploadInfo, p pipelineFunc, encrypt bool) (boson.Address, error) {
	// first store the file and get its reference
	fr, err := p(ctx, fileInfo.reader, fileInfo.size)
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("split file: %w", err)
	}

	// if filename is still empty, use the file hash as the filename
	if fileInfo.name == "" {
		fileInfo.name = fr.String()
	}

	// then store the metadata and get its reference
	m := entry.NewMetadata(fileInfo.name)
	m.MimeType = fileInfo.contentType
	metadataBytes, err := json.Marshal(m)
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("metadata marshal: %w", err)
	}


	mr, err := p(ctx, bytes.NewReader(metadataBytes), int64(len(metadataBytes)))
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("split metadata: %w", err)
	}

	// now join both references (mr, fr) to create an entry and store it
	e := entry.New(fr, mr)
	fileEntryBytes, err := e.MarshalBinary()
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("entry marshal: %w", err)
	}
	ref, err := p(ctx, bytes.NewReader(fileEntryBytes), int64(len(fileEntryBytes)))
	if err != nil {
		return boson.ZeroAddress, fmt.Errorf("split entry: %w", err)
	}

	return ref, nil
}
