package emboss

import (
	"context"
	"flag"
	"os"
	"testing"
)

const EXPECTED_REMOTE string = `Mood-lit Libations
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
"america`

var remote_embosser_uri = flag.String("remote-embosser-uri", "", "A valid sfomuseum/go-text-emboss URI")

func TestRemoteEmbosser(t *testing.T) {

	if *remote_embosser_uri == "" {
		t.Skip()
	}

	ctx := context.Background()

	e, err := NewEmbosser(ctx, *remote_embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	rsp, err := e.EmbossText(ctx, "fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := string(rsp)

	if str_rsp != EXPECTED_REMOTE {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}

func TestRemoteEmbosserWithReader(t *testing.T) {

	if *remote_embosser_uri == "" {
		t.Skip()
	}

	ctx := context.Background()

	e, err := NewEmbosser(ctx, *remote_embosser_uri)

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

	str_rsp := string(rsp)

	if str_rsp != EXPECTED_REMOTE {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}
