package handlers

import (
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/login"

	"github.com/gofiber/fiber/v2"
	middleware "github.com/katallaxie/fiber-htmx"
	htmx "github.com/katallaxie/htmx"
)

// UserLogin handles the user login page.
func UserLogin() middleware.CompFunc {
	return func(c *fiber.Ctx) (htmx.Node, error) {
		return components.Page(
			components.PageProps{},
			components.Wrap(
				components.WrapProps{},
				login.NewLogin(),
			),
		), nil
	}
}
