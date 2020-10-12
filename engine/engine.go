package engine

import (
	"github.com/dusansimic/feedgen/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Options stores options for the new engine instance
type Options struct {
	Origins []string
}

// New creates a new engine instance
func New(o Options) *gin.Engine {
	r := gin.Default()

	c := cors.DefaultConfig()
	c.AllowOrigins = o.Origins
	r.Use(cors.New(c))

	r.GET("/instagram/:user", handlers.InstagramHandler())
	r.GET("/informatikapmf", handlers.InformatikaPMFHandler())

	return r
}
