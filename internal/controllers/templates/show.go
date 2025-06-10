package templates

import (
	"context"

	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/templates"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	htmx "github.com/katallaxie/htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ShowTemplateControllerImpl ...
type ShowTemplateControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewShowTemplateController ...
func NewShowTemplateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowTemplateControllerImpl {
	return &ShowTemplateControllerImpl{
		store: store,
	}
}

// Get ...
func (l *ShowTemplateControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       "Template",
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
				template := models.Template{}

				errorx.Panic(l.BindParams(&template))
				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTemplate(ctx, &template)
				}))

				return htmx.Fragment(
					templates.TemplateTitleCard(
						templates.TemplateTitleCardProps{
							Template: template,
						},
					),
					templates.TemplateBodyCard(
						templates.TemplateBodyCardProps{
							Template: template,
						},
					),
				)
			},
		),
	)
}
