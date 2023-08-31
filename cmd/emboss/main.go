package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/sfomuseum/go-text-emboss"
)

func main() {

	embosser_uri := flag.String("embosser-uri", "local:///usr/local/sfomuseum/bin/text-emboss", "A valid sfomuseum/go-text-emboss.Embosser URI.")

	flag.Parse()

	ctx := context.Background()

	ex, err := emboss.NewEmbosser(ctx, *embosser_uri)

	if err != nil {
		log.Fatalf("Failed to create new embosser, %v", err)
	}

	for _, path := range flag.Args() {

		rsp, err := ex.EmbossText(ctx, path)

		if err != nil {
			log.Fatalf("Failed to extract text from %s, %v", path, err)
		}

		fmt.Println(string(rsp))
	}
}
