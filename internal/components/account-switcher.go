package components

import (
	"github.com/katallaxie/fiber-goth/v3/adapters"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/icons/heroicons"
)

// AccountSwitcherProps ...
type AccountSwitcherProps struct {
	// User ...
	User adapters.GothUser
	// ClassNames ...
	htmx.ClassNames
}

// AccountSwitcher ...
func AccountSwitcher(props AccountSwitcherProps, children ...htmx.Node) htmx.Node {
	return dropdowns.Dropdown(
		dropdowns.Props{},
		dropdowns.Button(
			dropdowns.ButtonProps{
				ClassNames: htmx.ClassNames{
					"btn":             true,
					"btn-sm":          true,
					"btn-outline":     true,
					"w-full":          true,
					"justify-between": true,
				},
			},
			htmx.Text("ZEISS"),
			heroicons.ChevronUpDownDefaultOutline(icons.IconProps{}),
		),
		dropdowns.Props(
			dropdowns.Props{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
		),
	)
}
