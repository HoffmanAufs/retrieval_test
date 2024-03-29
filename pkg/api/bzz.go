// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strings"

	"github.com/gauss-project/aurorafs/pkg/boson"
	"github.com/gauss-project/aurorafs/pkg/collection/entry"
	"github.com/gauss-project/aurorafs/pkg/file"
	"github.com/gauss-project/aurorafs/pkg/file/joiner"
	"github.com/gauss-project/aurorafs/pkg/file/loadsave"
	"github.com/gauss-project/aurorafs/pkg/jsonhttp"
	"github.com/gauss-project/aurorafs/pkg/manifest"
	"github.com/gauss-project/aurorafs/pkg/storage"
	"github.com/gauss-project/aurorafs/pkg/tracing"
)

func (s *server) bzzDownloadHandler(w http.ResponseWriter, r *http.Request) {
	logger := tracing.NewLoggerWithTraceID(r.Context(), s.logger)
	ls := loadsave.New(s.storer, storage.ModePutRequest, false)
	feedDereferenced := false


	ctx := r.Context()

	nameOrHex := mux.Vars(r)["address"]
	pathVar := mux.Vars(r)["path"]
	if strings.HasSuffix(pathVar, "/") {
		pathVar = strings.TrimRight(pathVar, "/")
		// NOTE: leave one slash if there was some
		pathVar += "/"
	}

	address, err := s.resolveNameOrAddress(nameOrHex)
	if err != nil {
		logger.Debugf("aurora download: parse address %s: %v", nameOrHex, err)
		logger.Error("aurora download: parse address")
		jsonhttp.NotFound(w, nil)
		return
	}

	if !s.chunkInfo.Init(r.Context(), nil, address) {
		logger.Debugf("aurora download: chunkInfo init %s: %v", nameOrHex, err)
		jsonhttp.NotFound(w, nil)
		return
	}

	// read manifest entry
	j, _, err := joiner.New(ctx, s.storer, address)
	if err != nil {
		logger.Debugf("aurora download: joiner manifest entry %s: %v", address, err)
		logger.Errorf("aurora download: joiner %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	buf := bytes.NewBuffer(nil)
	_, err = file.JoinReadAll(ctx, j, buf)
	if err != nil {
		logger.Debugf("aurora download: read entry %s: %v", address, err)
		logger.Errorf("aurora download: read entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	e := &entry.Entry{}
	err = e.UnmarshalBinary(buf.Bytes())
	if err != nil {
		logger.Debugf("aurora download: unmarshal entry %s: %v", address, err)
		logger.Errorf("aurora download: unmarshal entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	// read metadata
	j, _, err = joiner.New(ctx, s.storer, e.Metadata())
	if err != nil {
		logger.Debugf("aurora download: joiner metadata %s: %v", address, err)
		logger.Errorf("aurora download: joiner %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	// read metadata
	buf = bytes.NewBuffer(nil)
	_, err = file.JoinReadAll(ctx, j, buf)
	if err != nil {
		logger.Debugf("aurora download: read metadata %s: %v", address, err)
		logger.Errorf("aurora download: read metadata %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}
	manifestMetadata := &entry.Metadata{}
	err = json.Unmarshal(buf.Bytes(), manifestMetadata)
	if err != nil {
		logger.Debugf("aurora download: unmarshal metadata %s: %v", address, err)
		logger.Errorf("aurora download: unmarshal metadata %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	// we are expecting manifest Mime type here
	m, err := manifest.NewManifestReference(
		manifestMetadata.MimeType,
		e.Reference(),
		ls,
	)
	if err != nil {
		logger.Debugf("aurora download: not manifest %s: %v", address, err)
		logger.Error("aurora download: not manifest")
		jsonhttp.NotFound(w, nil)
		return
	}

	if pathVar == "" {
		logger.Tracef("aurora download: handle empty path %s", address)

		if indexDocumentSuffixKey, ok := manifestMetadataLoad(ctx, m, manifestRootPath, manifestWebsiteIndexDocumentSuffixKey); ok {
			pathWithIndex := path.Join(pathVar, indexDocumentSuffixKey)
			indexDocumentManifestEntry, err := m.Lookup(ctx, pathWithIndex)
			if err == nil {
				// index document exists
				logger.Debugf("aurora download: serving path: %s", pathWithIndex)

				s.serveManifestEntry(w, r,  address, indexDocumentManifestEntry.Reference(), !feedDereferenced)
				return
			}
		}
	}

	me, err := m.Lookup(ctx, pathVar)
	if err != nil {
		logger.Debugf("aurora download: invalid path %s/%s: %v", address, pathVar, err)
		logger.Error("aurora download: invalid path")

		if errors.Is(err, manifest.ErrNotFound) {

			if !strings.HasPrefix(pathVar, "/") {
				// check for directory
				dirPath := pathVar + "/"
				exists, err := m.HasPrefix(ctx, dirPath)
				if err == nil && exists {
					// redirect to directory
					u := r.URL
					u.Path += "/"
					redirectURL := u.String()

					logger.Debugf("aurora download: redirecting to %s: %v", redirectURL, err)

					http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
					return
				}
			}

			// check index suffix path
			if indexDocumentSuffixKey, ok := manifestMetadataLoad(ctx, m, manifestRootPath, manifestWebsiteIndexDocumentSuffixKey); ok {
				if !strings.HasSuffix(pathVar, indexDocumentSuffixKey) {
					// check if path is directory with index
					pathWithIndex := path.Join(pathVar, indexDocumentSuffixKey)
					indexDocumentManifestEntry, err := m.Lookup(ctx, pathWithIndex)
					if err == nil {
						// index document exists
						logger.Debugf("aurora download: serving path: %s", pathWithIndex)

						s.serveManifestEntry(w, r, address, indexDocumentManifestEntry.Reference(), !feedDereferenced)
						return
					}
				}
			}

			// check if error document is to be shown
			if errorDocumentPath, ok := manifestMetadataLoad(ctx, m, manifestRootPath, manifestWebsiteErrorDocumentPathKey); ok {
				if pathVar != errorDocumentPath {
					errorDocumentManifestEntry, err := m.Lookup(ctx, errorDocumentPath)
					if err == nil {
						// error document exists
						logger.Debugf("aurora download: serving path: %s", errorDocumentPath)

						s.serveManifestEntry(w, r, address, errorDocumentManifestEntry.Reference(), !feedDereferenced)
						return
					}
				}
			}

			jsonhttp.NotFound(w, "path address not found")
		} else {
			jsonhttp.NotFound(w, nil)
		}
		return
	}

	// serve requested path
	s.serveManifestEntry(w, r, address, me.Reference(), !feedDereferenced)
}

func (s *server) serveManifestEntry(w http.ResponseWriter, r *http.Request, address, manifestEntryAddress boson.Address, etag bool) {
	var (
		logger = tracing.NewLoggerWithTraceID(r.Context(), s.logger)
		ctx    = r.Context()
		buf    = bytes.NewBuffer(nil)
	)

	// read file entry
	j, _, err := joiner.New(ctx, s.storer, manifestEntryAddress)
	if err != nil {
		logger.Debugf("aurora download: joiner read file entry %s: %v", address, err)
		logger.Errorf("aurora download: joiner read file entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	_, err = file.JoinReadAll(ctx, j, buf)
	if err != nil {
		logger.Debugf("aurora download: read file entry %s: %v", address, err)
		logger.Errorf("aurora download: read file entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}
	fe := &entry.Entry{}
	err = fe.UnmarshalBinary(buf.Bytes())
	if err != nil {
		logger.Debugf("aurora download: unmarshal file entry %s: %v", address, err)
		logger.Errorf("aurora download: unmarshal file entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	// read file metadata
	j, _, err = joiner.New(ctx, s.storer, fe.Metadata())
	if err != nil {
		logger.Debugf("aurora download: joiner read file entry %s: %v", address, err)
		logger.Errorf("aurora download: joiner read file entry %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	buf = bytes.NewBuffer(nil)
	_, err = file.JoinReadAll(ctx, j, buf)
	if err != nil {
		logger.Debugf("aurora download: read file metadata %s: %v", address, err)
		logger.Errorf("aurora download: read file metadata %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}
	fileMetadata := &entry.Metadata{}
	err = json.Unmarshal(buf.Bytes(), fileMetadata)
	if err != nil {
		logger.Debugf("aurora download: unmarshal metadata %s: %v", address, err)
		logger.Errorf("aurora download: unmarshal metadata %s", address)
		jsonhttp.NotFound(w, nil)
		return
	}

	additionalHeaders := http.Header{
		"Content-Disposition": {fmt.Sprintf("inline; filename=\"%s\"", fileMetadata.Filename)},
		"Content-Type":        {fileMetadata.MimeType},
	}

	fileEntryAddress := fe.Reference()

	s.downloadHandler(w, r, fileEntryAddress, additionalHeaders, etag, address)
}

// manifestMetadataLoad returns the value for a key stored in the metadata of
// manifest path, or empty string if no value is present.
// The ok result indicates whether value was found in the metadata.
func manifestMetadataLoad(ctx context.Context, manifest manifest.Interface, path, metadataKey string) (string, bool) {
	me, err := manifest.Lookup(ctx, path)
	if err != nil {
		return "", false
	}

	manifestRootMetadata := me.Metadata()
	if val, ok := manifestRootMetadata[metadataKey]; ok {
		return val, ok
	}

	return "", false
}


