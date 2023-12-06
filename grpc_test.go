package emboss

import (
	"context"
	"flag"
	"os"
	"testing"
)

const EXPECTED_GRPC string = `Mood-lit Libations
Champagne Powder Cocktail
Champagne served with St. Germain
elderflower liqueur and hibiscus syrup
Mile-High Manhattan
Stranahans whiskey served with
sweet vermouth
Peach Collins On The Rockies
Silver Tree vodka, Leopold Bros peach
liqueur, lemon juice and agave nectar
Colorado Craft Beer
California Wines
america`

var grpc_embosser_uri = flag.String("grpc-embosser-uri", "", "A valid sfomuseum/go-text-emboss URI")

func TestGRPCEmbosser(t *testing.T) {

	if *grpc_embosser_uri == "" {
		t.Skip()
	}

	ctx := context.Background()

	e, err := NewEmbosser(ctx, *grpc_embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	rsp, err := e.EmbossText(ctx, "fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := rsp.String()

	if str_rsp != EXPECTED_GRPC {
		t.Fatalf("Unexpected output '%s', expected '%s'", str_rsp, EXPECTED_GRPC)
	}
}

func TestGRPCEmbosserWithReader(t *testing.T) {

	if *grpc_embosser_uri == "" {
		t.Skip()
	}

	ctx := context.Background()

	e, err := NewEmbosser(ctx, *grpc_embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	r, err := os.Open("fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to open test image, %v", err)
	}

	defer r.Close()

	rsp, err := e.EmbossTextWithReader(ctx, "fixtures/menu.jpg", r)

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := rsp.String()

	if str_rsp != EXPECTED_GRPC {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}
