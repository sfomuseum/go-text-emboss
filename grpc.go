package emboss

import (
	"context"
	"fmt"
	"io"
	_ "log"
	"net/url"
	"os"
	"path/filepath"

	emboss_grpc "github.com/sfomuseum/go-text-emboss/grpc"
	"google.golang.org/grpc"
)

type GrpcEmbosser struct {
	endpoint string
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "grpc", NewGrpcEmbosser)
}

func NewGrpcEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	addr := u.Host

	e := &GrpcEmbosser{
		endpoint: addr,
	}

	return e, nil
}

func (e *GrpcEmbosser) EmbossText(ctx context.Context, path string) ([]byte, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossTextWithReader(ctx, path, im_r)
}

func (e *GrpcEmbosser) EmbossTextWithReader(ctx context.Context, path string, im_r io.Reader) ([]byte, error) {

	fname := filepath.Base(path)

	body, err := io.ReadAll(im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", path, err)
	}

	// START OF grpc conn stuff

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(e.endpoint, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial '%s', %w", e.endpoint, err)
	}

	defer conn.Close()

	// END OF grpc conn stuff

	req := &emboss_grpc.EmbossTextRequest{
		Filename: fname,
		Body:     body,
	}

	client := emboss_grpc.NewEmbosserClient(conn)

	rsp, err := client.EmbossText(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to emboss text, %w", err)
	}

	return []byte(rsp.Body), nil
}
