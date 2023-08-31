# go-text-emboss

Go package for interacting with the `sfomuseum/swift-text-emboss` command-line and www tools

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

### Remote

Remote text "embossing" depends on their being a copy of the [text-emboss-server](https://github.com/sfomuseum/swift-text-emboss-www) tool running that can be accessed over HTTP.

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

## Tests

To run tests you will need to specify the custom `-local-embosser-uri` and `-remote-embosser-uri` flag with values specific to your system. For example:

```
$> go test -v -local-embosser-uri local:///usr/local/sfomuseum/bin/text-emboss -remote-embosser-uri http://localhost:8080
=== RUN   TestLocalEmbosser
--- PASS: TestLocalEmbosser (0.34s)
=== RUN   TestLocalEmbosserWithReader
--- PASS: TestLocalEmbosserWithReader (0.29s)
=== RUN   TestLocalEmbosserWithReaderAndPath
--- PASS: TestLocalEmbosserWithReaderAndPath (0.29s)
=== RUN   TestNullEmbosser
--- PASS: TestNullEmbosser (0.00s)
=== RUN   TestNullEmbosserWithReader
--- PASS: TestNullEmbosserWithReader (0.00s)
=== RUN   TestRemoteEmbosser
--- PASS: TestRemoteEmbosser (0.30s)
=== RUN   TestRemoteEmbosserWithReader
--- PASS: TestRemoteEmbosserWithReader (0.29s)
PASS
ok
```

_Note that the `TestLocal` tests are only applicable on OS X (`darwin`) systems._

## See also

* https://github.com/sfomuseum/swift-text-emboss
* https://github.com/sfomuseum/swift-text-emboss-cli
* https://github.com/sfomuseum/swift-text-emboss-www
* https://collection.sfomuseum.org/objects/1880246621/