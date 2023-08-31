package emboss

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

type LocalEmbosser struct {
	text_cli string
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "local", NewLocalEmbosser)
}

func NewLocalEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	text_cli := u.Path

	_, err = os.Stat(text_cli)

	if err != nil {
		return nil, fmt.Errorf("Failed to stat %s, %w", text_cli, err)
	}

	e := &LocalEmbosser{
		text_cli: text_cli,
	}

	return e, nil
}

func (e *LocalEmbosser) EmbossText(ctx context.Context, path string) ([]byte, error) {

	args := []string{
		path,
	}

	out, err := exec.CommandContext(ctx, e.text_cli, args...).Output()

	if err != nil {
		return nil, fmt.Errorf("Failed to extract text, %w", err)
	}

	return out, nil
}
