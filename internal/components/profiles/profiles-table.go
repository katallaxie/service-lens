package profiles

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

// ProfilesTableProps ...
type ProfilesTableProps struct {
	URL      string
	Profiles []*models.Profile
	Team     string
	Offset   int
	Limit    int
	Total    int
}

// ProfilesTable ...
func ProfilesTable(props ProfilesTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.Props{
				ID: "profiles-tables",
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
						htmx.Href(utils.CreateProfileUrlFormat),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Text("Add Profile"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.Profile]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.Props) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.Props, row *models.Profile) htmx.Node {
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
					Cell: func(p tables.Props, row *models.Profile) htmx.Node {
						return htmx.Td(
							links.Link(
								links.Props{
									Href: fmt.Sprintf(utils.ShowProfileUrlFormat, row.ID.String()),
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
					Cell: func(p tables.Props, row *models.Profile) htmx.Node {
						return htmx.Td(
							buttons.Button(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"btn-sm": true,
									},
								},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteProfileUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete this profile?"),
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
			props.Profiles,
		),
	)
}
