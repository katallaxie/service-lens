package environments

import (
	"fmt"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/icons/heroicons"
	"github.com/katallaxie/htmx/links"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"
)

// EnvironmentsTableProps ...
type EnvironmentsTableProps struct {
	URL          string
	Environments []*models.Environment
	Offset       int
	Limit        int
	Total        int
}

// EnvironmentsTable ...
func EnvironmentsTable(props EnvironmentsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.Props{
				ID: "environments-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{},
					tables.Pagination(
						tables.PaginationProps{},
						tables.Prev(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),

						tables.Select(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								Limits: tables.DefaultLimits,
								URL:    props.URL,
							},
						),
						tables.Next(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),
					),
				),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"inline-flex":  true,
							"items-center": true,
							"gap-3":        true,
						},
						tables.Search(
							tables.SearchProps{
								Name:        "search",
								Placeholder: "Search ...",
								URL:         props.URL,
							},
						),
					),
					htmx.A(
						htmx.Href("/environments/new"),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Text("Create Environment"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Environment]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.Props) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.Props, row *models.Environment) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.Props) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.Props, row *models.Environment) htmx.Node {
						return htmx.Td(
							links.Link(
								links.Props{
									Href: "/environments/" + row.ID.String(),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.Props) htmx.Node {
						return nil
					},
					Cell: func(p tables.Props, row *models.Environment) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteEnvironmentUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete environment tag?"),
								htmx.HxTarget("closest tr"),
								htmx.HxSwap("outerHTML swap:1s"),
								heroicons.TrashDefaultOutline(
									icons.IconProps{
										ClassNames: htmx.ClassNames{
											"w-6 h-6": false,
											"w-4":     true,
											"h-4":     true,
										},
									},
								),
							),
						)
					},
				},
			},
			props.Environments,
		),
	)
}
