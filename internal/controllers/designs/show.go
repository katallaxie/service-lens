package designs

import (
	"context"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/designs"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	htmx "github.com/katallaxie/htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ShowDesignControllerImpl ...
type ShowDesignControllerImpl struct {
	design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// NewShowDesignController ...
func NewShowDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowDesignControllerImpl {
	return &ShowDesignControllerImpl{
		store: store,
	}
}

// Clone ...
func (i *ShowDesignControllerImpl) Clone() handlers.Controller {
	return &ShowDesignControllerImpl{store: i.store}
}

// Prepare ...
func (l *ShowDesignControllerImpl) Prepare() error {
	err := l.BindAll(&l.design)
	if err != nil {
		return err
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.design)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *ShowDesignControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       l.design.Title,
				Path:        l.Ctx().Path(),
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
				return htmx.Fragment(
					designs.DesignTitleCard(
						designs.DesignTitleCardProps{
							Design: l.design,
						},
					),
					designs.DesignBodyCard(
						designs.DesignBodyCardProps{
							User:   l.Session().User,
							Design: l.design,
						},
					),
					designs.DesignMetadataCard(
						designs.DesignMetadataCardProps{
							Design: l.design,
						},
					),
					designs.DesignRevisionsCard(
						designs.DesignRevisionsCardProps{
							DesignID: l.design.ID,
						},
					),
					designs.DesignTagsCard(
						designs.DesignTagsCardProps{
							Design: l.design,
						},
					),
					designs.DesignCommentsCard(
						designs.DesignCommentsCardProps{
							User:   l.Session().User,
							Design: l.design,
						},
					),
				)
			},
		),
	)
}
