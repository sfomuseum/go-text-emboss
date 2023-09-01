package emboss

import (
	"context"
	"fmt"
	"io"
	_ "log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"google.golang.org/grpc"	
)

type GrpcEmbosser struct {
	client   *http.Client
	endpoint string
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "grpc", NewGrpcEmbosser)
}

func NewGrpcEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse()

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

	// START OF grpc stuff

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(e.endpoint, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial '%s', %v", addr, err)
	}

	defer conn.Close()

	req := &foo.Request{
		Filename: fname,
		Body: body,
	}
	
	client := foo.NewClient(conn)
	
	// END OF grpc stuff
	
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status '%s'", rsp.Status)
	}

	rsp_body, err := io.ReadAll(rsp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read response, %w", err)
	}

	return rsp_body, nil
}
