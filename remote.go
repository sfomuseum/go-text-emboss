package emboss

import (
	"bytes"
	"context"
	"fmt"
	"io"
	_ "log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
)

type RemoteEmbosser struct {
	client   *http.Client
	endpoint string
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "http", NewRemoteEmbosser)
	RegisterEmbosser(ctx, "https", NewRemoteEmbosser)
}

func NewRemoteEmbosser(ctx context.Context, uri string) (Embosser, error) {

	cl := &http.Client{}

	e := &RemoteEmbosser{
		client:   cl,
		endpoint: uri,
	}

	return e, nil
}

func (e *RemoteEmbosser) EmbossText(ctx context.Context, path string) ([]byte, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossTextWithReader(ctx, path, im_r)
}

func (e *RemoteEmbosser) EmbossTextWithReader(ctx context.Context, path string, im_r io.Reader) ([]byte, error) {

	fname := filepath.Base(path)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("Content-Type", "image/jpeg")

	part, err := CreateImageFormFile(writer, "image", fname)

	if err != nil {
		return nil, fmt.Errorf("Failed to create form part, %w", err)
	}

	n, err := io.Copy(part, im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to copy image to form, %w", err)
	}

	err = writer.Close()

	if err != nil {
		return nil, fmt.Errorf("Failed to close form writer, %w", err)
	}

	req, err := http.NewRequest("POST", e.endpoint, body)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request, %w", err)
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("ContentLength", strconv.FormatInt(n, 10))

	req = req.WithContext(ctx)

	rsp, err := e.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute request, %w", err)
	}

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

func CreateImageFormFile(w *multipart.Writer, name, filename string) (io.Writer, error) {

	ext := filepath.Ext(filename)
	content_type := mime.TypeByExtension(ext)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, name, filename))
	h.Set("Content-Type", content_type)
	return w.CreatePart(h)
}
