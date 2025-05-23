package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfomuseum/go-text-emboss/v2"
)

func main() {

	embosser_uri := flag.String("embosser-uri", "local:///usr/local/sfomuseum/bin/text-emboss", "A valid sfomuseum/go-text-emboss.Embosser URI.")
	as_json := flag.Bool("as-json", false, "Return results as a JSON-encoded dictionary containing text, source and creation time properties.")

	flag.Parse()

	ctx := context.Background()

	em, err := emboss.NewEmbosser(ctx, *embosser_uri)

	if err != nil {
		log.Fatalf("Failed to create new embosser, %v", err)
	}

	defer em.Close(ctx)

	for _, path := range flag.Args() {

		rsp, err := em.EmbossText(ctx, path)

		if err != nil {
			log.Fatalf("Failed to extract text from %s, %v", path, err)
		}

		if *as_json {

			enc := json.NewEncoder(os.Stdout)
			err = enc.Encode(rsp)

			if err != nil {
				log.Fatalf("Failed to encode response as JSON, %v", err)
			}

		} else {
			fmt.Println(rsp.String())
		}
	}
}
