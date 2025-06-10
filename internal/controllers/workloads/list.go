package workloads

import (
	"context"

	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/workloads"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
)

// WorkloadListControllerImpl ...
type WorkloadListControllerImpl struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadListController ...
func NewWorkloadListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadListControllerImpl {
	return &WorkloadListControllerImpl{store: store}
}

// Get ...
func (w *WorkloadListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				results := tables.Results[models.Workload]{SearchFields: []string{"name"}}

				errorx.Panic(w.BindQuery(&results))
				errorx.Panic(w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.ListWorkloads(ctx, &results)
				}))

				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						workloads.WorkloadsTable(
							workloads.WorkloadsTableProps{
								Workloads: results.GetRows(),
								Offset:    results.GetOffset(),
								Limit:     results.GetLimit(),
								Total:     results.GetTotalRows(),
								URL:       w.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
