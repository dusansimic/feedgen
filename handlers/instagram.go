package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dusansimic/feedgen"
	grab "github.com/dusansimic/feedgen/internal/grab/instagram"
	source "github.com/dusansimic/feedgen/internal/source/instagram"
	"github.com/gin-gonic/gin"
)

var failResponseFunc = func(ctx *gin.Context) {
	ctx.Status(500)
	ctx.Abort()
}

func InstagramHandler(clientID, clientSecret string) gin.HandlerFunc {
	url := fmt.Sprintf("https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret)
	resp, err := http.Get(url)
	if err != nil {
		return failResponseFunc
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var tokenData struct {
		Token string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		return failResponseFunc
	}

	return func(ctx *gin.Context) {
		username := ctx.Param("user")

		var q queries
		if err := ctx.BindQuery(&q); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
			ctx.Abort()
			return
		}

		g := grab.New(grab.WithSource(source.S{
			Username: username,
			Token:    tokenData.Token,
			Last:     q.Last,
		}))

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
