package components

import (
	"strings"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/menus"
	"github.com/katallaxie/service-lens/internal/utils"
)

// MainMenuProps ...
type MainMenuProps struct {
	Path string
	Team string
	htmx.ClassNames
}

// MainMenu ...
func MainMenu(p MainMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		menus.Menu(
			menus.Props{
				ClassNames: htmx.ClassNames{
					"min-h-full": true,
					"w-80":       true,
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
			menus.Title(
				menus.TitleProps{},
				htmx.Text("Design & Review"),
			),
			menus.Item(
				menus.ItemProps{
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
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
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
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
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
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
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
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
					ClassNames: htmx.ClassNames{
						"hover:bg-base-300": false,
					},
				},
				menus.Link(
					menus.LinkProps{
						Href:   utils.ListTemplatesUrlFormat,
						Active: strings.HasPrefix(p.Path, utils.ListTemplatesUrlFormat),
					},
					htmx.Text("Templates"),
				),
			),
		),
	)
}
