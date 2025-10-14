package tags

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/tags"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/dbx"
	seed "github.com/zeiss/gorm-seed"
)

// IndexController ...
type IndexController struct {
	model dbx.Results[models.Tag]
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
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

// Prepare ...
func (i *IndexController) Prepare() error {
	if err := i.BindQuery(&i.model); err != nil {
		return err
	}

	err := i.store.ReadTx(i.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTags(ctx, &i.model)
	})
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (i *IndexController) Get() error {
	return i.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        i.Path(),
				User:        i.Session().User,
				Development: i.IsDevelopment(),
			},
			func() htmx.Node {
				return cards.CardBorder(
					cards.Props{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						tags.TagsTable(
							tags.TagsTableProps{
								Tags:   i.model.GetRows(),
								Offset: i.model.GetOffset(),
								Limit:  i.model.GetLimit(),
								Total:  i.model.GetTotalRows(),
								URL:    i.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
