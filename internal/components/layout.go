package components

import (
	middleware "github.com/katallaxie/fiber-htmx"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dividers"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/icons/heroicons"
	"github.com/katallaxie/htmx/navbars"
	"github.com/zeiss/fiber-goth/adapters"
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
	return htmx.Div(
		htmx.ClassNames{},
		htmx.Div(
			htmx.ClassNames{},
			drawers.Drawer(
				drawers.DrawerProps{
					ID: "global-drawer",
					ClassNames: htmx.ClassNames{
						"lg:drawer-open": true,
					},
				},
				drawers.DrawerContent(
					drawers.DrawerContentProps{
						ID: "drawer",
					},
					middleware.Toasts(),
					htmx.Div(
						htmx.ClassNames{
							"h-full":        true,
							"max-w-full":    true,
							"overflow-auto": true,
							"w-full":        true,
						},
						htmx.Div(
							htmx.ClassNames{
								"flex":        true,
								"h-full":      true,
								"flex-col":    true,
								"bg-base-200": true,
							},
							navbars.Navbar(
								navbars.NavbarProps{
									ClassNames: htmx.ClassNames{
										"navbar":      true,
										"z-10":        true,
										"px-3":        true,
										"bg-base-100": true,
									},
								},
								navbars.NavbarStart(
									navbars.NavbarStartProps{
										ClassNames: htmx.ClassNames{
											"gap-3": true,
										},
									},
									drawers.DrawerOpenButton(
										drawers.DrawerOpenProps{
											ID: "global-drawer",
											ClassNames: htmx.ClassNames{
												"btn-sm":      true,
												"btn-square":  true,
												"btn-primary": false,
											},
										},
										heroicons.Bars3Outline(
											icons.IconProps{},
										),
									),
								),
								navbars.NavbarEnd(
									navbars.NavbarEndProps{},
									ProfileMenu(
										ProfileMenuProps{
											User: p.User,
										},
									),
								),
							),
							htmx.Group(
								children...,
							),
						),
					),
				),
				drawers.DrawerSide(
					drawers.DrawerSideProps{
						ID: "drawer",
					},
					drawers.DrawerSideMenu(
						drawers.DrawerSideMenuProps{
							ClassNames: htmx.ClassNames{
								"border-r":               true,
								"border-neutral-content": true,
								"bg-base-100":            true,
								"bg-base-200":            false,
							},
						},
						dividers.Divider(
							dividers.DividerProps{},
						),
						MainMenu(
							MainMenuProps{
								Path: p.Path,
							},
						),
						dividers.Divider(
							dividers.DividerProps{},
						),
						UserMenu(
							UserMenuProps{
								Path: p.Path,
							},
						),
					),
				),
			),
		),
	)
}
