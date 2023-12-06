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

func (e *NullEmbosser) EmbossText(ctx context.Context, path string) (*EmbossTextResult, error) {
	return e.nullEmbossTextResult(), nil
}

func (e *NullEmbosser) EmbossTextWithReader(ctx context.Context, path string, r io.Reader) (*EmbossTextResult, error) {
	return e.nullEmbossTextResult(), nil
}

func (e *NullEmbosser) Close(ctx context.Context) error {
	return nil
}

func (e *NullEmbosser) nullEmbossTextResult() *EmbossTextResult {

	r := &EmbossTextResult{
		Text:    "",
		Source:  "",
		Created: 0,
	}

	return r
}
