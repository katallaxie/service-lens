package components

import (
	"strings"

	"github.com/katallaxie/fiber-goth/v3/adapters"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dividers"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/menus"
	"github.com/katallaxie/htmx/navbars"
	"github.com/katallaxie/service-lens/internal/utils"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	// Team is the teams to user adapters.
	User adapters.GothUser
	Path string
}

// WrapProps ...
type WrapProps struct {
	ClassNames htmx.ClassNames
}

// Wrap ...
func Wrap(p WrapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Layout is a whole document to output.
func Layout(p LayoutProps, children ...htmx.Node) htmx.Node {
	return drawers.Drawer(
		drawers.Props{
			Open: true,
		},
		drawers.Toggle(
			drawers.ToggleProps{
				ID: "app-drawer",
			},
		),
		drawers.Content(
			drawers.ContentProps{},
			navbars.Navbar(
				navbars.Props{
					ClassNames: htmx.ClassNames{
						"w-full":      true,
						"bg-base-300": true,
					},
				},
				navbars.Start(
					navbars.StartProps{},
					htmx.Label(
						htmx.For("app-drawer"),
						htmx.ClassNames{
							"btn":        true,
							"btn-ghost":  true,
							"btn-circle": true,
							"lg:hidden":  true,
						},
						htmx.Span(
							htmx.ClassNames{
								"material-symbols-outlined": true,
							},
							htmx.Text("menu"),
						),
					),
				),
				navbars.End(
					navbars.EndProps{},
					ProfileMenu(
						ProfileMenuProps{
							User: p.User,
						},
					),
				),
			),
			Wrap(
				WrapProps{
					ClassNames: htmx.ClassNames{
						"p-4":         true,
						"min-h-full":  true,
						"bg-base-100": true,
					},
				},
				children...,
			),
		),
		drawers.Side(
			drawers.SideProps{
				ClassNames: htmx.ClassNames{
					"w-80":            true,
					"bg-base-200":     true,
					"min-h-full":      true,
					"flex":            true,
					"flex-col":        true,
					"justify-between": true,
				},
			},
			menus.Menu(
				menus.Props{
					ClassNames: htmx.ClassNames{
						"bg-base-200": true,
						"rounded-box": true,
						"w-80":        true,
						"p-4":         true,
					},
				},
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href:   "/",
							Active: p.Path == "/",
						},
						htmx.Text("Dashboard"),
					),
				),
				dividers.Divider(
					dividers.Props{},
				),
				menus.Title(
					menus.TitleProps{},
					htmx.Text("Design & Review"),
				),
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href:   "/designs",
							Active: strings.HasPrefix(p.Path, "/designs"),
						},
						htmx.Text("Designs"),
					),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{},
					},
					menus.Link(
						menus.LinkProps{
							Href:   "/workloads",
							Active: strings.HasPrefix(p.Path, "/workloads"),
						},
						htmx.Text("Workloads"),
					),
				),
				menus.Title(
					menus.TitleProps{},
					htmx.Text("Configuration"),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.Link(
						menus.LinkProps{
							Href:   "/lenses",
							Active: strings.HasPrefix(p.Path, "/lenses"),
						},
						htmx.Text("Lenses"),
					),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.Link(
						menus.LinkProps{
							Href:   "/profiles",
							Active: strings.HasPrefix(p.Path, "/profiles"),
						},
						htmx.Text("Profiles"),
					),
				),
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href:   "/environments",
							Active: strings.HasPrefix(p.Path, "/environments"),
						},
						htmx.Text("Environments"),
					),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{},
					},
					menus.Link(
						menus.LinkProps{
							Href:   utils.ListTagsUrlFormat,
							Active: strings.HasPrefix(p.Path, "/tags"),
						},
						htmx.Text("Tags"),
					),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{},
					},
					menus.Link(
						menus.LinkProps{
							Href:   utils.ListWorkflowsUrlFormat,
							Active: strings.HasPrefix(p.Path, "/workflows"),
						},
						htmx.Text("Workflows"),
					),
				),
				menus.Item(
					menus.ItemProps{
						ClassNames: htmx.ClassNames{},
					},
					menus.Link(
						menus.LinkProps{
							Href:   utils.ListTemplatesUrlFormat,
							Active: strings.HasPrefix(p.Path, utils.ListTemplatesUrlFormat),
						},
						htmx.Text("Templates"),
					),
				),
				dividers.Divider(
					dividers.Props{},
				),
				menus.Title(
					menus.TitleProps{},
					htmx.Text("Settings"),
				),
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href:   "/settings",
							Active: strings.HasPrefix(p.Path, "/settings"),
						},
						htmx.Text("Settings"),
					),
				),
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href:   "/me",
							Active: strings.HasPrefix(p.Path, "/me"),
						},
						htmx.Text("Profile"),
					),
				),
				menus.Item(
					menus.ItemProps{},
					menus.Link(
						menus.LinkProps{
							Href: "/logout",
						},
						htmx.Text("Logout"),
					),
				),
			),
		),
	)
}
