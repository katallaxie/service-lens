package templates

import (
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/templates"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
)

// NewTemplateControllerImpl ...
type NewTemplateControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTemplateController ...
func NewTemplateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *NewTemplateControllerImpl {
	return &NewTemplateControllerImpl{store: store}
}

// Get ...
func (l *NewTemplateControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.css"),
						htmx.Rel("stylesheet"),
						htmx.Type("text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://cdn.jsdelivr.net/simplemde/1.11/simplemde.min.js"),
						htmx.Type("text/javascript"),
					),
				},
			},
			func() htmx.Node {
				return templates.TemplateNewForm(
					templates.TemplateNewFormProps{},
				)
			},
		),
	)
}
