package components

import (
	"github.com/katallaxie/fiber-goth/v3/adapters"
	middleware "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/menus"
	"github.com/katallaxie/htmx/navbars"
	"github.com/katallaxie/htmx/typography"
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
			ClassNames: htmx.ClassNames{
				"drawer-open": true,
			},
		},
		drawers.Toggle(
			drawers.ToggleProps{
				ID: "global-drawer",
			},
		),
		drawers.Content(
			drawers.ContentProps{},
			middleware.Toasts(),
			navbars.Navbar(
				navbars.Props{},
				htmx.Div(
					htmx.ClassNames{
						"mx-2":   true,
						"flex-1": true,
						"px-2":   true,
						"py-1":   true,
					},
					typography.H4(
						typography.Props{},
						htmx.Text("Service Lens"),
					),
				),
				navbars.End(
					navbars.EndProps{},
					menus.MenuHorizontal(
						menus.Props{},
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
						menus.Item(
							menus.ItemProps{},
							menus.Link(
								menus.LinkProps{
									Href:   "/designs",
									Active: p.Path == "/designs",
								},
								htmx.Text("Designs"),
							),
						),
						menus.Item(
							menus.ItemProps{},
							menus.Link(
								menus.LinkProps{
									Href:   "/lenses",
									Active: p.Path == "/lenses",
								},
								htmx.Text("Lenses"),
							),
						),
						menus.Item(
							menus.ItemProps{},
							menus.Link(
								menus.LinkProps{
									Href:   "/workloads",
									Active: p.Path == "/workloads",
								},
								htmx.Text("Workloads"),
							),
						),
						menus.Item(
							menus.ItemProps{},
							menus.Link(
								menus.LinkProps{
									Href:   "/settings",
									Active: p.Path == "/settings",
								},
								htmx.Text("Settings"),
							),
						),
					),
					ProfileMenu(
						ProfileMenuProps{
							User: p.User,
						},
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"h-full": true,
					"w-full": true,
				},
				htmx.Group(
					children...,
				),
			),
			// htmx.Div(
			// 	htmx.ClassNames{
			// 		"h-full":        true,
			// 		"max-w-full":    true,
			// 		"overflow-auto": true,
			// 		"w-full":        true,
			// 	},
			// 	htmx.Div(
			// 		htmx.ClassNames{
			// 			"flex":        true,
			// 			"h-full":      true,
			// 			"flex-col":    true,
			// 			"bg-base-200": true,
			// 		},
			// 		navbars.Navbar(
			// 			navbars.Props{
			// 				ClassNames: htmx.ClassNames{
			// 					"navbar":      true,
			// 					"z-10":        true,
			// 					"px-3":        true,
			// 					"bg-base-100": true,
			// 				},
			// 			},
			// 			navbars.Start(
			// 				navbars.StartProps{
			// 					ClassNames: htmx.ClassNames{
			// 						"gap-3": true,
			// 					},
			// 				},
			// 				drawers.Toggle(
			// 					drawers.ToggleProps{
			// 						ID: "global-drawer",
			// 						ClassNames: htmx.ClassNames{
			// 							"btn-sm":      true,
			// 							"btn-square":  true,
			// 							"btn-primary": false,
			// 						},
			// 					},
			// 					heroicons.Bars3Outline(
			// 						icons.IconProps{},
			// 					),
			// 				),
			// 			),
			// 			navbars.End(
			// 				navbars.EndProps{},
			// 				ProfileMenu(
			// 					ProfileMenuProps{
			// 						User: p.User,
			// 					},
			// 				),
			// 			),
			// 		),
			// 		htmx.Group(
			// 			children...,
			// 		),
			// 	),
			// ),
		),
		// drawers.Side(
		// 	drawers.SideProps{
		// 		ClassNames: htmx.ClassNames{
		// 			"border-r":               true,
		// 			"border-neutral-content": true,
		// 			"bg-base-100":            true,
		// 			"bg-base-200":            false,
		// 		},
		// 	},
		// 	MainMenu(
		// 		MainMenuProps{
		// 			Path: p.Path,
		// 		},
		// 	),
		// 	dividers.Divider(
		// 		dividers.Props{},
		// 	),
		// 	UserMenu(
		// 		UserMenuProps{
		// 			Path: p.Path,
		// 		},
		// 	),
		// ),
	)
}
