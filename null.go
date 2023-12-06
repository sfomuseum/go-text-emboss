package emboss

import (
	"context"
	"io"
)

type NullEmbosser struct {
	text_cli string
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "null", NewNullEmbosser)
}

func NewNullEmbosser(ctx context.Context, uri string) (Embosser, error) {
	e := &NullEmbosser{}
	return e, nil
}

func (e *NullEmbosser) EmbossText(ctx context.Context, path string) ([]byte, error) {
	return []byte(""), nil
}

func (e *NullEmbosser) EmbossTextWithReader(ctx context.Context, path string, r io.Reader) ([]byte, error) {
	return []byte(""), nil
}

func (e *NullEmbosser) EmbossTextAsResult(ctx context.Context, path string) (*ProcessImageResult, error) {
	return e.nullProcessImageResult(), nil
}

func (e *NullEmbosser) EmbossTextAsResultWithReader(ctx context.Context, path string, r io.Reader) (*ProcessImageResult, error) {
	return e.nullProcessImageResult(), nil
}

func (e *NullEmbosser) Close(ctx context.Context) error {
	return nil
}

func (e *NullEmbosser) nullProcessImageResult() *ProcessImageResult {

	r := &ProcessImageResult{
		Text:    "",
		Source:  "",
		Created: 0,
	}

	return r
}
