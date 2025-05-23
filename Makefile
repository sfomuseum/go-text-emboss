GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/emboss cmd/emboss/main.go

# https://developers.google.com/protocol-buffers/docs/reference/go-generated
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#
# https://github.com/sfomuseum/swift-text-emboss-grpc/blob/main/Sources/text-emboss-grpc-server/embosser.proto

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/org_sfomuseum_text_embosser.proto

test:
	go run -mod $(GOMOD) cmd/emboss/main.go -embosser-uri grpc://localhost:8080 ./fixtures/menu.jpg
