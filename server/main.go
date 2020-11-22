package main

import (
	"os"
	"strconv"

	"github.com/dusansimic/feedgen/engine"
)

func main() {
	enableInstagram, _ := strconv.ParseBool(os.Getenv("ENABLE_INSTAGRAM"))
	instagramClientID := os.Getenv("INSTAGRAM_CLIENT_ID")
	instagramClientSecret := os.Getenv("INSTAGRAM_CLIENT_SECRET")

	e := engine.New(
		engine.WithAllowedOrigin("*"),
		engine.WithParameter("ENABLE_INSTAGRAM", enableInstagram),
		engine.WithParameter("INSTAGRAM_CLIENT_ID", instagramClientID),
		engine.WithParameter("INSTAGRAM_CLIENT_SECRET", instagramClientSecret),
	)

	e.Run(":3000")
}
