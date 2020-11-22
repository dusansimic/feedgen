package engine

import (
	"github.com/dusansimic/feedgen/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Options stores options for the new engine instance
type options struct {
	Origins []string
	Params  map[string]interface{}
}

// Option sets an option
type Option func(*options)

// New creates a new engine instance
func New(opts ...Option) *gin.Engine {
	// Set options
	o := options{}
	o.Params = make(map[string]interface{})
	for _, opt := range opts {
		opt(&o)
	}

	// Create a new router
	r := gin.Default()

	// Set cors settings
	c := cors.DefaultConfig()
	c.AllowOrigins = o.Origins
	r.Use(cors.New(c))

	// If instagram is enabled, it requires client id and secret for thumbnails
	if o.Params["ENABLE_INSTAGRAM"].(bool) {
		r.GET("/instagram/:user", handlers.InstagramHandler(o.Params["INSTAGRAM_CLIENT_ID"].(string), o.Params["INSTAGRAM_CLIENT_SECRET"].(string)))
	}

	r.GET("/informatikapmf", handlers.InformatikaPMFHandler())

	return r
}

// WithAllowedOrigin sets an allowed origin
func WithAllowedOrigin(org string) Option {
	return func(o *options) {
		o.Origins = append(o.Origins, org)
	}
}

// WithParameter sets a parameter for later use in grabbers
func WithParameter(k string, v interface{}) Option {
	return func(o *options) {
		o.Params[k] = v
	}
}
