package main

import "github.com/go-humble/router"

func stories(context *router.Context) {
	print("stories")
	fadeIn("jass-model-cycle")
	noButtons()
}
