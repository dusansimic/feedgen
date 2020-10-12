package instagram

import (
	"fmt"
	"time"

	"github.com/dusansimic/feedgen"
	content "github.com/dusansimic/feedgen/internal/content/instagram"
)

// New creates a new Instagram feed with passed options.
func New(opts ...Option) feedgen.Feeder {
	f := &F{}
	for _, o := range opts {
		o(f)
	}
	return f
}

// F stores information about Instagram feed.
type F struct {
	FullName    string
	Bio         string
	Username    string
	Profile     string
	UserID      string
	UpdatedTime int
	Content     []content.C
}

// Title returns the title of the feed.
func (f F) Title() string {
	return f.FullName
}

// Subtitle returns the subtitle of the feed.
func (f F) Subtitle() string {
	return f.Bio
}

// Icon returns the icon url for the feed.
func (f F) Icon() string {
	return f.Profile
}

// Links returns the urls for the feed.
func (f F) Links() []feedgen.Link {
	links := []feedgen.Link{
		{
			Href: fmt.Sprintf("https://instagram.com/%s/", f.Username),
			Rel:  "self",
		},
	}
	return links
}

// ID returns the id of the feed.
func (f F) ID() string {
	return f.UserID
}

// Updated retunrs the last time feed was updated.
func (f F) Updated() time.Time {
	return time.Unix(int64(f.UpdatedTime), 0)
}

// Entries returns feed entries.
func (f F) Entries() []feedgen.Contenter {
	cs := make([]feedgen.Contenter, len(f.Content))
	for i, e := range f.Content {
		cs[i] = feedgen.Contenter(e)
	}
	return cs
}
