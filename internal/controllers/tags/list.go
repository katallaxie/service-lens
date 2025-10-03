package tags

import (
	"context"

	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/tags"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
	seed "github.com/zeiss/gorm-seed"
)

// TagsListControllerImpl ...
type TagsListControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagsListController ...
func NewTagsListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TagsListControllerImpl {
	return &TagsListControllerImpl{store: store}
}

// Get ...
func (c *TagsListControllerImpl) Get() error {
	return c.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        c.Path(),
				User:        c.Session().User,
				Development: c.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Tag]{SearchFields: []string{"Name"}}

				errorx.Panic(c.BindQuery(&results))
				errorx.Panic(c.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListTags(ctx, &results)
				}))

				return cards.CardBordered(
					cards.Props{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						tags.TagsTable(
							tags.TagsTableProps{
								Tags:   results.GetRows(),
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
