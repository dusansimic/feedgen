package instagram

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dusansimic/feedgen"
	content "github.com/dusansimic/feedgen/internal/content/instagram"
	feed "github.com/dusansimic/feedgen/internal/feed/instagram"
	source "github.com/dusansimic/feedgen/internal/source/instagram"
	"github.com/gocolly/colly/v2"
)

// New creates a new grabber for Instagram with specified options.
func New(opts ...Option) feedgen.Grabber {
	g := &G{}
	for _, o := range opts {
		o(g)
	}
	return g
}

// G is a grabber implementation for Instagram.
type G struct {
	Source source.S
}

// Grab creates a Feeder instance and populates it with data from Instagram
func (g *G) Grab() feedgen.Feeder {
	data, err := scrapeData(g.Source.Username)
	if err != nil {
		fmt.Println(err)
	}

	user := data.EntryData.ProfilePage[0].GraphQL.User

	cs := make([]content.C, len(user.Media.Edges))
	for i, e := range user.Media.Edges {
		var c content.C

		c.Username = g.Source.Username
		c.Caption = e.Node.Caption.Edges[0].Node.Text
		c.MediaID = e.Node.Shortcode
		c.ThumbURL = getThumbURL(g.Source.Token, e.Node.Shortcode)
		c.Medias = make([]string, len(e.Node.Children.Edges))
		c.Thumbs = make([]string, len(e.Node.Children.Edges))
		for i, e := range e.Node.Children.Edges {
			c.Medias[i] = e.Node.Shortcode
			c.Thumbs[i] = getThumbURL(g.Source.Token, e.Node.Shortcode)
		}
		c.Time = e.Node.Date

		cs[i] = c
	}

	f := feed.New(
		feed.WithUser(user.FullName, user.Biography, user.Username, user.ProfileURL, user.ID),
		feed.WithContent(cs),
		feed.WithUpdated(getMaxTime(cs)),
	)

	return f
}

func scrapeData(u string) (data *DataJSON, err error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", fmt.Sprintf("https://www.instagram.com/%s", u))
		if r.Ctx.Get("gis") != "" {
			gis := fmt.Sprintf("%s:%s", r.Ctx.Get("gis"), r.Ctx.Get("variables"))
			h := md5.New()
			h.Write([]byte(gis))
			gisHash := fmt.Sprintf("%x", h.Sum(nil))
			r.Headers.Set("X-Instagram-GIS", gisHash)
		}
	})

	c.OnHTML("html", func(h *colly.HTMLElement) {
		dat := h.ChildText("body > script:first-of-type")
		jsonData := dat[strings.Index(dat, "{") : len(dat)-1]
		data = &DataJSON{}
		cbErr := json.Unmarshal([]byte(jsonData), data)
		if cbErr != nil {
			err = fmt.Errorf("can't unmarshal data: %w", cbErr)
			return
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		err = fmt.Errorf("error on request '%s %s': %w", r.Request.Method, r.Request.URL, e)
	})

	c.Visit(fmt.Sprintf("https://www.instagram.com/%s", u))
	c.Wait()

	return data, err
}

func getMaxTime(c []content.C) int {
	var max int
	for i, cont := range c {
		if i == 0 || max < cont.Time {
			max = cont.Time
		}
	}
	return max
}

func getThumbURL(tkn, sc string) string {
	// Create a new default client
	client := http.DefaultClient
	// Create a new get request for instagram pic
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://graph.facebook.com/v9.0/instagram_oembed?url=https://www.instagram.com/p/%s/", sc), nil)
	// Add authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tkn))
	// Execute the request
	resp, _ := client.Do(req)
	// Close the body after reading
	defer resp.Body.Close()
	// Read data from body
	body, _ := ioutil.ReadAll(resp.Body)
	// Unmarshal body
	var data oembedData
	_ = json.Unmarshal(body, &data)
	// Return the thumb url
	return data.ThumbURL
}
