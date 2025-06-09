package components

import (
	"strings"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/menus"
)

// UserMenuProps ...
type UserMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
}

// UserMenu ...
func UserMenu(p UserMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
		),
		menus.Menu(
			menus.Props{
				ClassNames: htmx.ClassNames{
					"w-full":      true,
					"bg-base-200": false,
				},
			},
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
	)
}
