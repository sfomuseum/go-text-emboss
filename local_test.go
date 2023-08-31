package emboss

import (
	"context"
	"os"
	"testing"
)

const EXPECTED_LOCAL string = `Mood-lit Libations
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
"america
`

func TestLocalEmbosser(t *testing.T) {

	ctx := context.Background()

	embosser_uri := "local:///usr/local/sfomuseum/bin/text-emboss"

	e, err := NewEmbosser(ctx, embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	rsp, err := e.EmbossText(ctx, "fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := string(rsp)

	if str_rsp != EXPECTED_LOCAL {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}

func TestLocalEmbosserWithReader(t *testing.T) {

	ctx := context.Background()

	embosser_uri := "local:///usr/local/sfomuseum/bin/text-emboss"

	e, err := NewEmbosser(ctx, embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	r, err := os.Open("fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to open test image, %v", err)
	}

	defer r.Close()

	rsp, err := e.EmbossTextWithReader(ctx, "", r)

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := string(rsp)

	if str_rsp != EXPECTED_LOCAL {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}

func TestLocalEmbosserWithReaderAndPath(t *testing.T) {

	ctx := context.Background()

	embosser_uri := "local:///usr/local/sfomuseum/bin/text-emboss"

	e, err := NewEmbosser(ctx, embosser_uri)

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	r, err := os.Open("fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to open test image, %v", err)
	}

	defer r.Close()

	rsp, err := e.EmbossTextWithReader(ctx, "fixtures/test.jpg", r)

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := string(rsp)

	if str_rsp != EXPECTED_LOCAL {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}
