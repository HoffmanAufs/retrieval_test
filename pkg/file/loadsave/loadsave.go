package loadsave

import (
	"bytes"
	"context"

	"github.com/gauss-project/aurorafs/pkg/file"
	"github.com/gauss-project/aurorafs/pkg/file/joiner"
	"github.com/gauss-project/aurorafs/pkg/file/pipeline/builder"
	"github.com/gauss-project/aurorafs/pkg/storage"
	"github.com/gauss-project/aurorafs/pkg/boson"
)

// loadSave is needed for manifest operations and provides
// simple wrapping over load and save operations using file
// package abstractions. use with caution since Loader will
// load all of the subtrie of a given hash in memory.
type loadSave struct {
	storer    storage.Storer
	mode      storage.ModePut
	encrypted bool
}

func New(storer storage.Storer, mode storage.ModePut, enc bool) file.LoadSaver {
	return &loadSave{
		storer:    storer,
		mode:      mode,
		encrypted: enc,
	}
}

func (ls *loadSave) Load(ctx context.Context, ref []byte) ([]byte, error) {
	j, _, err := joiner.New(ctx, ls.storer, boson.NewAddress(ref))
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = file.JoinReadAll(ctx, j, buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (ls *loadSave) Save(ctx context.Context, data []byte) ([]byte, error) {
	pipe := builder.NewPipelineBuilder(ctx, ls.storer, ls.mode, ls.encrypted)
	address, err := builder.FeedPipeline(ctx, pipe, bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return boson.ZeroAddress.Bytes(), err
	}

	return address.Bytes(), nil

}
