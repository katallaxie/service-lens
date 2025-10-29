package lenses

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/lenses"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/pkg/dbx"
	seed "github.com/zeiss/gorm-seed"
)

// ListController ...
type ListController struct {
	results dbx.Results[models.Lens]
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// Clone ...
func (l *ListController) Clone() handlers.Controller {
	return &ListController{store: l.store}
}

// NewListController ...
func NewListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListController {
	return &ListController{store: store}
}

func (l *ListController) Prepare() error {
	if err := l.BindQuery(&l.results); err != nil {
		return err
	}

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListLenses(ctx, &l.results)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *ListController) Get() error {
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
						lenses.LensesTable(
							lenses.LensesTableProps{
								Lenses: l.results.GetRows(),
								Offset: l.results.GetOffset(),
								Limit:  l.results.GetLimit(),
								Total:  l.results.GetLen(),
								URL:    l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
