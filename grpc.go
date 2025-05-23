package emboss

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	_ "log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	emboss_grpc "github.com/sfomuseum/go-text-emboss/v2/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcEmbosser struct {
	conn   *grpc.ClientConn
	client emboss_grpc.TextEmbosserClient
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

	q := u.Query()

	q_tls_cert := q.Get("tls-certificate")
	q_tls_key := q.Get("tls-key")
	q_tls_ca := q.Get("tls-ca-certificate")
	q_tls_insecure := q.Get("tls-insecure")

	addr := u.Host

	opts := make([]grpc.DialOption, 0)

	if q_tls_cert != "" && q_tls_key != "" {

		cert, err := tls.LoadX509KeyPair(q_tls_cert, q_tls_key)

		if err != nil {
			return nil, fmt.Errorf("Failed to load TLS pair, %w", err)
		}

		tls_config := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		if q_tls_ca != "" {

			ca_cert, err := ioutil.ReadFile(q_tls_ca)

			if err != nil {
				return nil, fmt.Errorf("Failed to create CA certificate, %w", err)
			}

			cert_pool := x509.NewCertPool()

			ok := cert_pool.AppendCertsFromPEM(ca_cert)

			if !ok {
				return nil, fmt.Errorf("Failed to append CA certificate, %w", err)
			}

			tls_config.RootCAs = cert_pool

		} else if q_tls_insecure != "" {

			v, err := strconv.ParseBool(q_tls_insecure)

			if err != nil {
				return nil, fmt.Errorf("Failed to parse ?tls-insecure= parameter, %w", err)
			}

			tls_config.InsecureSkipVerify = v
		}

		tls_credentials := credentials.NewTLS(tls_config)
		opts = append(opts, grpc.WithTransportCredentials(tls_credentials))

	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial '%s', %w", addr, err)
	}

	// defer conn.Close()

	client := emboss_grpc.NewTextEmbosserClient(conn)

	e := &GrpcEmbosser{
		conn:   conn,
		client: client,
	}

	return e, nil
}

func (e *GrpcEmbosser) EmbossText(ctx context.Context, path string) (*EmbossTextResult, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossTextWithReader(ctx, path, im_r)
}

func (e *GrpcEmbosser) EmbossTextWithReader(ctx context.Context, path string, im_r io.Reader) (*EmbossTextResult, error) {

	fname := filepath.Base(path)

	body, err := io.ReadAll(im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", path, err)
	}

	req := &emboss_grpc.EmbossTextRequest{
		Filename: fname,
		Body:     body,
	}

	pb_rsp, err := e.client.EmbossText(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to emboss text, %w", err)
	}

	rsp := &EmbossTextResult{
		Text:    string(pb_rsp.Result.Text),
		Source:  pb_rsp.Result.Source,
		Created: pb_rsp.Result.Created,
	}

	return rsp, nil
}

func (e *GrpcEmbosser) Close(ctx context.Context) error {
	return e.conn.Close()
}
