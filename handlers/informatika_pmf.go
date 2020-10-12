package handlers

import (
	"net/http"
	"time"

	"github.com/dusansimic/feedgen"
	grab "github.com/dusansimic/feedgen/internal/grab/informatika_pmf"
	"github.com/gin-gonic/gin"
)

// InformatikaPMFHandler handles Informatika PMF feed requests
func InformatikaPMFHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g := grab.New()
		f := g.Grab()

		content := f.Entries()

		var entries []feedgen.Entry
		for _, e := range content {
			var links []feedgen.Link
			for _, l := range e.Links() {
				links = append(links, feedgen.Link{
					Href: l,
				})
			}

			entries = append(entries, feedgen.Entry{
				Title:   e.Title(),
				Links:   links,
				ID:      e.ID(),
				Updated: e.Updated().Format(time.RFC3339),
				Summary: e.Summary(),
				Content: feedgen.Content{
					Type: "html",
					Body: e.Content(),
				},
				Author: feedgen.Author{
					Name: e.Author(),
				},
			})
		}

		image := feedgen.Image{
			Title: f.Title(),
			URL:   f.Icon(),
			Link:  f.Links()[0].Href,
		}

		feed := feedgen.Feed{
			Title:    f.Title(),
			Subtitle: f.Subtitle(),
			Image:    image,
			Links:    f.Links(),
			ID:       f.ID(),
			Updated:  f.Updated().Format(time.RFC3339),
			Entries:  entries,
		}

		ctx.XML(http.StatusOK, feed)
	}
}
