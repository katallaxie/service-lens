package components

import (
	"github.com/katallaxie/fiber-goth/adapters"
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
		dropdowns.DropdownProps{},
		dropdowns.DropdownButton(
			dropdowns.DropdownButtonProps{
				ClassNames: htmx.ClassNames{
					"btn":             true,
					"btn-sm":          true,
					"btn-outline":     true,
					"w-full":          true,
					"justify-between": true,
				},
			},
			htmx.Text("ZEISS"),
			heroicons.ChevronUpDownOutline(icons.IconProps{}),
		),
		dropdowns.DropdownMenuItems(
			dropdowns.DropdownMenuItemsProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
		),
	)
}
