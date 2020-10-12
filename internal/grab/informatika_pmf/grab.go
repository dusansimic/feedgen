package informatikapmf

import (
	"fmt"
	"time"

	"github.com/dusansimic/feedgen"
	content "github.com/dusansimic/feedgen/internal/content/informatika_pmf"
	feed "github.com/dusansimic/feedgen/internal/feed/informatika_pmf"
	"github.com/gocolly/colly/v2"
)

// New creates a new grabber for Informatika PMF with specified options.
func New() feedgen.Grabber {
	return &G{}
}

// G is a grabber implementation for Informatika PMF
type G struct {
}

// Grab returns grabbed feed
func (g *G) Grab() feedgen.Feeder {
	c := colly.NewCollector()

	var contents []content.C

	c.OnHTML("article.post", func(h *colly.HTMLElement) {
		var c content.C

		titleEl := h.DOM.Find("h2.entry-title")
		c.EntryTitle = titleEl.Text()
		c.URL = titleEl.Find("a").AttrOr("href", "")

		timeEl := h.DOM.Find("time")
		c.Time = timeEl.AttrOr("datetime", time.Now().Format(time.RFC3339))

		contentEl := h.DOM.Find("div.entry-content")
		contentHTML, err := contentEl.Html()
		if err != nil {
			fmt.Println(err)
		}
		c.EntryContent = contentHTML

		footerEl := h.DOM.Find("footer.entry-footer")
		footerHTML, err := footerEl.Html()
		if err != nil {
			fmt.Println(err)
		}
		c.EntryContent += footerHTML

		categoryEl := footerEl.Find("a")
		c.Category = categoryEl.Text()

		contents = append(contents, c)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("colly error: %s", e.Error())
	})

	c.Visit("https://informatika.pmf.uns.ac.rs/vesti/")
	c.Wait()

	f := feed.New(
		feed.WithDefaults(),
		feed.WithUpdated(contents[0].Updated()),
		feed.WithContent(contents),
	)

	return f
}
