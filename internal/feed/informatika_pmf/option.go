package informatikapmf

import (
	"time"

	content "github.com/dusansimic/feedgen/internal/content/informatika_pmf"
)

// Option sets an option for feed
type Option func(*F)

// WithDefaults sets default options for feed
func WithDefaults() Option {
	return func(f *F) {
		f.Name = "Informatika PMF"
		f.URL = "https://informatika.pmf.uns.ac.rs"
		f.IconURL = "https://informatika.pmf.uns.ac.rs/wp-content/uploads/2020/02/cropped-oi-logo-1-2-192x192.png"
		f.UpdatedTime = time.Now()
	}
}

// WithUpdated sets updated time to passed time
func WithUpdated(t time.Time) Option {
	return func(f *F) {
		f.UpdatedTime = t
	}
}

// WithContent sets the content to passed content
func WithContent(c []content.C) Option {
	return func(f *F) {
		f.Content = c
	}
}
