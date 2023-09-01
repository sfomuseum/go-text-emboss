# go-text-emboss

Go package for interacting with the `sfomuseum/swift-text-emboss` command-line, www and grpc tools

## Documentation

Documentation is incomplete.

## Example

_Error handling omitted for the sake of brevity._

### Local

Local text "embossing" depends on their being a copy of the [text-emboss](https://github.com/sfomuseum/swift-text-emboss-cli) tool that can be invoked in a shell command.

```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sfomuseum/go-text-emboss"
)

func main() {

	ctx := context.Background()
	ex, _ := emboss.NewEmbosser(ctx, "local:///usr/local/sfomuseum/bin/text-emboss")

	for _, path := range flag.Args() {
		rsp, _ := ex.EmbossText(ctx, path)
		fmt.Println(string(rsp))
	}
}
```

### Remote (HTTP)

Remote (HTTP) text "embossing" depends on their being a copy of the [text-emboss-server](https://github.com/sfomuseum/swift-text-emboss-www) tool running that can be accessed over HTTP.

```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sfomuseum/go-text-emboss"
)

func main() {

	ctx := context.Background()
	ex, _ := emboss.NewEmbosser(ctx, "http://localhost:8080")

	for _, path := range flag.Args() {
		rsp, _ := ex.EmbossText(ctx, path)
		fmt.Println(string(rsp))
	}
}
```

### Remote (gRPC)

Remote (gRPC) text "embossing" depends on their being a copy of the [text-emboss-server](https://github.com/sfomuseum/swift-text-emboss-grpc) tool running that can be accessed over TCP.

```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sfomuseum/go-text-emboss"
)

func main() {

	ctx := context.Background()
	ex, _ := emboss.NewEmbosser(ctx, "grpc://localhost:1234")

	for _, path := range flag.Args() {
		rsp, _ := ex.EmbossText(ctx, path)
		fmt.Println(string(rsp))
	}
}
```

## Tools

### emboss

```
$> ./bin/emboss -h
Usage of ./bin/emboss:
  -embosser-uri string
    	A valid sfomuseum/go-text-emboss.Embosser URI. (default "local:///usr/local/sfomuseum/bin/text-emboss")
```

For example, assuming there is a copy of the [text-emboss](https://github.com/sfomuseum/swift-text-emboss-cli) tool at `/usr/local/sfomuseum/bin/text-emboss`:

```
$> ./bin/emboss .fixtures/menu.jpg 
Mood-lit Libations
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
```

Or, assuming the there is a copy of the [text-emboss-server](https://github.com/sfomuseum/swift-text-emboss-www) tool listening for requests on `http://localhost:8080`:

```
$> ./bin/emboss -embosser-uri http://localhost:8080 .fixtures/menu.jpg
Mood-lit Libations
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
```

Or, assuming the there is a copy of the [text-emboss-server](https://github.com/sfomuseum/swift-text-emboss-grpc) tool listening for requests on `localhost:1234`:

```
$> ./bin/emboss -embosser-uri grpc://localhost:1234 .fixtures/menu.jpg
Mood-lit Libations
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
```

## Tests

To run tests you will need to specify the custom `-local-embosser-uri`, `-grpc-embosser-uri` and `-http-embosser-uri` flag with values specific to your system. For example:

```
$> go test -v -local-embosser-uri local:///usr/local/sfomuseum/bin/text-emboss -http-embosser-uri http://localhost:8080 -grpc-embosser-uri grpc://localhost:1234
=== RUN   TestGRPCEmbosser
--- PASS: TestGRPCEmbosser (0.09s)
=== RUN   TestGRPCEmbosserWithReader
--- PASS: TestGRPCEmbosserWithReader (0.07s)
=== RUN   TestHTTPEmbosser
--- PASS: TestHTTPEmbosser (0.73s)
=== RUN   TestHTTPEmbosserWithReader
--- PASS: TestHTTPEmbosserWithReader (0.85s)
=== RUN   TestLocalEmbosser
--- PASS: TestLocalEmbosser (0.31s)
=== RUN   TestLocalEmbosserWithReader
--- PASS: TestLocalEmbosserWithReader (0.29s)
=== RUN   TestLocalEmbosserWithReaderAndPath
--- PASS: TestLocalEmbosserWithReaderAndPath (0.30s)
=== RUN   TestNullEmbosser
--- PASS: TestNullEmbosser (0.00s)
=== RUN   TestNullEmbosserWithReader
--- PASS: TestNullEmbosserWithReader (0.00s)
PASS
ok  	github.com/sfomuseum/go-text-emboss	2.791s
```

_Note that the `TestLocal` tests are only applicable on OS X (`darwin`) systems._

## See also

* https://github.com/sfomuseum/swift-text-emboss
* https://github.com/sfomuseum/swift-text-emboss-cli
* https://github.com/sfomuseum/swift-text-emboss-www
* https://github.com/sfomuseum/swift-text-emboss-grpc
* https://collection.sfomuseum.org/objects/1880246621/