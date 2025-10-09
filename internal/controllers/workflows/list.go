package workflows

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/workflows"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
	seed "github.com/zeiss/gorm-seed"
)

// ListController ...
type ListController struct {
	model tables.Results[models.Workflow]
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// Clone ...
func (i *ListController) Clone() handlers.Controller {
	return &ListController{store: i.store}
}

// NewListController ...
func NewListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ListController {
	return &ListController{store: store}
}

// Prepare ...
func (i *ListController) Prepare() error {
	if err := i.BindQuery(&i.model); err != nil {
		return err
	}

	err := i.store.ReadTx(i.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListWorkflows(ctx, &i.model)
	})
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (i *ListController) Get() error {
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
						workflows.WorkflowsTable(
							workflows.WorkflowsTableProps{
								Workflows: i.model.GetRows(),
								Offset:    i.model.GetOffset(),
								Limit:     i.model.GetLimit(),
								Total:     i.model.GetTotalRows(),
								URL:       i.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
