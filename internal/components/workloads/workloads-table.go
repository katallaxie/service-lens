package workloads

import (
	"fmt"

	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/links"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
)

const (
	profileShowURL     = "/profiles/%s"
	environmentShowURL = "/environments/%s"
	deleteWorkloadURL  = "/workloads/%s"
)

// WorkloadsTableProps ...
type WorkloadsTableProps struct {
	URL       string
	Workloads []*models.Workload
	Offset    int
	Limit     int
	Total     int
}

// WorkloadsTable ...
func WorkloadsTable(props WorkloadsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "workloads-tables",
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
							tables.SelectProps{
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
						htmx.Href(utils.CreateWorkloadUrlFormat),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Text("Add Workload"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Workload]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: "/workloads/" + row.ID.String(),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "profile",
					AccessorKey: "profile",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Profile"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(profileShowURL, row.Profile.ID),
								},
								htmx.Text(row.Profile.Name),
							),
						)
					},
				},
				{
					ID:          "environment",
					AccessorKey: "environment",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Environment"))
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(environmentShowURL, row.Environment.ID),
								},
								htmx.Text(row.Environment.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Workload) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteWorkloadUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete workload?"),
								htmx.HxTarget("closest tr"),
								htmx.HxSwap("outerHTML swap:1s"),
								icons.TrashOutline(
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
			props.Workloads,
		),
	)
}
