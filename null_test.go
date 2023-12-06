package emboss

import (
	"context"
	"os"
	"testing"
)

const EXPECTED_NULL string = ""

func TestNullEmbosser(t *testing.T) {

	ctx := context.Background()

	e, err := NewEmbosser(ctx, "null://")

	if err != nil {
		t.Fatalf("Failed to create embosser, %v", err)
	}

	rsp, err := e.EmbossText(ctx, "fixtures/menu.jpg")

	if err != nil {
		t.Fatalf("Failed to emboss text, %v", err)
	}

	str_rsp := rsp.String()

	if str_rsp != EXPECTED_NULL {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}

func TestNullEmbosserWithReader(t *testing.T) {

	ctx := context.Background()

	e, err := NewEmbosser(ctx, "null://")

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

	if str_rsp != EXPECTED_NULL {
		t.Fatalf("Unexpected output '%s'", str_rsp)
	}
}
