package emboss

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-roster"
)

type Embosser interface {
	EmbossText(context.Context, string) ([]byte, error)
	EmbossTextWithReader(context.Context, string, io.Reader) ([]byte, error)
}

var embosser_roster roster.Roster

// EmbosserInitializationFunc is a function defined by individual embosser package and used to create
// an instance of that embosser
type EmbosserInitializationFunc func(ctx context.Context, uri string) (Embosser, error)

// RegisterEmbosser registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Embosser` instances by the `NewEmbosser` method.
func RegisterEmbosser(ctx context.Context, scheme string, init_func EmbosserInitializationFunc) error {

	err := ensureEmbosserRoster()

	if err != nil {
		return err
	}

	return embosser_roster.Register(ctx, scheme, init_func)
}

func ensureEmbosserRoster() error {

	if embosser_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		embosser_roster = r
	}

	return nil
}

// NewEmbosser returns a new `Embosser` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `EmbosserInitializationFunc`
// function used to instantiate the new `Embosser`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterEmbosser` method.
func NewEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := embosser_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(EmbosserInitializationFunc)
	return init_func(ctx, uri)
}

// Schemes returns the list of schemes that have been registered.
func Schemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureEmbosserRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range embosser_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
