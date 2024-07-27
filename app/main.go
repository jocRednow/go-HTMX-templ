package main

import (
	"github.com/jocRednow/go-HTMX-templ"
	"github.com/jocRednow/go-HTMX-templ/app/views/profile/"
)

func main() {
	app := fast.New()

	app.Get("/profile", HandleUserProfile)

	app.Start(":3000")
}

func HandleUserProfile(c *fast.Context) error {
	return c.Render(profile.Index())
}
