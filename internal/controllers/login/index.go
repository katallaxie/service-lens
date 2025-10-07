package login

import (
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/login"

	htmx "github.com/katallaxie/fiber-htmx/v3"
)

// IndexLoginController ...
type IndexLoginController struct {
	htmx.UnimplementedController
}

func (c *IndexLoginController) Clone() htmx.Controller {
	return &IndexLoginController{}
}

// NewIndexLoginController ...
func NewIndexLoginController() *IndexLoginController {
	return &IndexLoginController{}
}

// Get ...
func (l *IndexLoginController) Get() error {
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
