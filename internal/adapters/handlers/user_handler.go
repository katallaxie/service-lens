package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/login"
	htmx "github.com/zeiss/fiber-htmx"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Login() htmx.CompFunc {
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
