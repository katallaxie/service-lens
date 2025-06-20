package workloads

import (
	"github.com/katallaxie/service-lens/internal/models"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/conv"
)

// WorkloadMetadataCardProps ...
type WorkloadMetadataCardProps struct {
	Workload models.Workload
}

// WorkloadMetadataCard ...
func WorkloadMetadataCard(props WorkloadMetadataCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.ClassNames{
				tailwind.M2: true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Metadata"),
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
					htmx.Text("ID"),
				),
				htmx.H3(htmx.Text(conv.String(props.Workload.ID))),
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
						htmx.Text("Created at"),
					),
					htmx.H3(
						htmx.Text(
							props.Workload.CreatedAt.Format("2006-01-02 15:04:05"),
						),
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
						htmx.Text("Updated at"),
					),
					htmx.H3(
						htmx.Text(
							props.Workload.UpdatedAt.Format("2006-01-02 15:04:05"),
						),
					),
				),
			),
		),
	)
}
