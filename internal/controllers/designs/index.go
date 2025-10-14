package designs

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/designs"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/pkg/dbx"
	seed "github.com/zeiss/gorm-seed"
)

// IndexController ...
type IndexController struct {
	results dbx.Results[models.Design]
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// Clone ...
func (i *IndexController) Clone() handlers.Controller {
	return &IndexController{store: i.store}
}

// NewIndexController ...
func NewIndexController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *IndexController {
	return &IndexController{store: store}
}

func (l *IndexController) Prepare() error {
	if err := l.BindQuery(&l.results); err != nil {
		return err
	}

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListDesigns(ctx, &l.results)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *IndexController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				return cards.CardBorder(
					cards.Props{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						designs.DesignsTable(
							designs.DesignsTableProps{
								Designs: l.results.GetRows(),
								Offset:  l.results.GetOffset(),
								Limit:   l.results.GetLimit(),
								Total:   l.results.GetLen(),
								Search:  l.results.GetSearch(),
								URL:     l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
