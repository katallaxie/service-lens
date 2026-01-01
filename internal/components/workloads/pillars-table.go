package workloads

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/links"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/service-lens/internal/models"
)

// LensPillarTableProps ...
type LensPillarTableProps struct {
	Lens   *models.Lens
	Offset int
	Limit  int
	Total  int
}

// LensPillarTable ...
func LensPillarTable(props LensPillarTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.Props{
				ID: "lenses-pillar-table",
			},
			[]tables.ColumnDef[*models.Pillar]{
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.Props) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.Props, row *models.Pillar) htmx.Node {
						return htmx.Td(
							links.Link(
								links.Props{
									// Href: fmt.Sprintf(workloadLensURL, props.Workload.ID, row.ID),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
			},
			props.Lens.GetPillars(),
		),
	)
}
