package emboss

import (
	"context"
	"fmt"
	"io"
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

func (e *LocalEmbosser) EmbossTextWithReader(ctx context.Context, path string, r io.Reader) ([]byte, error) {

	var wr io.WriteCloser

	if path != "" {

		w, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			return nil, fmt.Errorf("Failed to open %s for writing, %w", path, err)
		}

		wr = w

	} else {

		w, err := os.CreateTemp("", "emboss")

		if err != nil {
			return nil, fmt.Errorf("Failed to create temp file for writing reader, %w", err)
		}

		path = w.Name()
		wr = w
	}

	defer os.Remove(path)

	_, err := io.Copy(wr, r)

	if err != nil {
		return nil, fmt.Errorf("Failed to copy reader to %s, %w", path, err)
	}

	err = wr.Close()

	if err != nil {
		return nil, fmt.Errorf("Failed to close %s, %w", path, err)
	}

	return e.EmbossText(ctx, path)
}
