package login

import (
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/login"

	htmx "github.com/katallaxie/fiber-htmx/v3"
)

// IndexController ...
type IndexController struct {
	htmx.UnimplementedController
}

func (c *IndexController) Clone() htmx.Controller {
	return &IndexController{}
}

// NewIndexLoginController ...
func NewIndexLoginController() *IndexController {
	return &IndexController{}
}

// Get ...
func (l *IndexController) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			components.Wrap(
				components.WrapProps{},
				login.NewLogin(),
			),
		),
	)
}
