package main

import "github.com/dusansimic/feedgen/engine"

func main() {
	e := engine.New(engine.Options{
		Origins: []string{
			"*",
		},
	})

	e.Run(":3000")
}
