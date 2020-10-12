package informatikapmf

import (
	"time"

	"github.com/dusansimic/feedgen"
	content "github.com/dusansimic/feedgen/internal/content/informatika_pmf"
)

// New creates a new Informatika PMF feed with passed options.
func New(opts ...Option) feedgen.Feeder {
	f := &F{}
	for _, o := range opts {
		o(f)
	}
	return f
}

// F stores info about a feed.
type F struct {
	Name        string
	IconURL     string
	URL         string
	UpdatedTime time.Time
	Content     []content.C
}

// Title returns the title of the feed.
func (f F) Title() string {
	return f.Name
}

// Subtitle returns the subtitle of the feed.
func (f F) Subtitle() string {
	return ""
}

// Icon returns the icon url of the feed.
func (f F) Icon() string {
	return f.IconURL
}

// Links returns the links for the feed.
func (f F) Links() []feedgen.Link {
	links := []feedgen.Link{
		{
			Href: f.URL,
			Rel:  "self",
		},
	}
	return links
}

// ID returns id of the feed.
func (f F) ID() string {
	return f.Name
}

// Updated returns time when the feed was last updated.
func (f F) Updated() time.Time {
	return f.UpdatedTime
}

// Entries reutrns the feed entries.
func (f F) Entries() []feedgen.Contenter {
	cs := make([]feedgen.Contenter, len(f.Content))
	for i, e := range f.Content {
		cs[i] = feedgen.Contenter(e)
	}
	return cs
}
