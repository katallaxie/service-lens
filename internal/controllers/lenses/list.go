package lenses

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/lenses"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/errorx"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
)

// LensListController ...
type LensListController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensListController ...
func NewLensListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensListController {
	return &LensListController{store: store}
}

// Get ...
func (c *LensListController) Get() error {
	return c.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        c.Path(),
				User:        c.Session().User,
				Development: c.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Lens]{SearchFields: []string{"Name"}}

				errorx.Panic(c.BindQuery(&results))
				errorx.Panic(c.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListLenses(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						lenses.LensesTable(
							lenses.LensesTableProps{
								Lenses: results.GetRows(),
								Offset: results.GetOffset(),
								Limit:  results.GetLimit(),
								Total:  results.GetTotalRows(),
								URL:    c.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
