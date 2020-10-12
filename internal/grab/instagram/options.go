package instagram

import source "github.com/dusansimic/feedgen/internal/source/instagram"

// Option sets a option for the Grabber
type Option func(*G)

// WithSource sets a source for Grabber
func WithSource(s source.S) Option {
	return func(g *G) {
		g.Source = s
	}
}
