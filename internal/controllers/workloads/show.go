package workloads

import (
	"context"
	"fmt"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/workloads"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
)

// WorkloadShowControllerImpl ...
type WorkloadShowControllerImpl struct {
	workload models.Workload
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadShowController ...
func NewWorkloadShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadShowControllerImpl {
	return &WorkloadShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *WorkloadShowControllerImpl) Prepare() error {
	err := w.BindParams(&w.workload)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetWorkload(ctx, &w.workload)
	})
}

// Get ...
func (w *WorkloadShowControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title:       w.workload.Name,
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				return htmx.Fragment(
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						htmx.HxTarget("this"),
						htmx.HxSwap("outerHTML"),
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Name"),
								),
								htmx.H3(
									htmx.Text(w.workload.Name),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Description"),
								),
								htmx.H3(
									htmx.Text(w.workload.Description),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{
										Type: "button",
									},
									htmx.Text("Edit"),
									htmx.HxGet(fmt.Sprintf(utils.EditWorkloadUrlFormat, w.workload.ID)),
									htmx.HxSwap("outerHTML"),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this workload?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
					workloads.WorkloadsRisksCard(
						workloads.WorkloadsRisksCardProps{
							Workload: w.workload,
						},
					),
					workloads.WorkloadMetadataCard(
						workloads.WorkloadMetadataCardProps{
							Workload: w.workload,
						},
					),
					workloads.WorkloadTagsCard(
						workloads.WorkloadTagsCardProps{
							Workload: w.workload,
						},
					),
					workloads.WorkloadProfileCard(
						workloads.WorkloadProfileCardProps{
							Workload: w.workload,
						},
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Lenses"),
							),
							workloads.LensesTable(
								workloads.LensesTableProps{
									Workload: w.workload,
								},
							),
						),
					),
				)
			},
		),
	)
}
