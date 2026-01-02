package components

import (
	"github.com/katallaxie/fiber-goth/v3/adapters"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/avatars"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/katallaxie/pkg/cast"
)

// ProfileMenuProps ...
type ProfileMenuProps struct {
	ClassNames htmx.ClassNames
	User       adapters.GothUser
}

// ProfileMenu ...
func ProfileMenu(p ProfileMenuProps, children ...htmx.Node) htmx.Node {
	return dropdowns.Dropdown(
		dropdowns.Props{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"dropdown-end": true,
				},
				p.ClassNames,
			),
		},
		dropdowns.Button(
			dropdowns.ButtonProps{
				ClassNames: htmx.ClassNames{
					"m-1":        true,
					"btn-circle": true,
					"btn-ghost":  true,
				},
			},
			avatars.RoundSmall(
				avatars.Props{},
				htmx.Img(
					htmx.Attribute("src", cast.Value(p.User.Image)),
				),
			),
		),
		dropdowns.MenuItems(
			dropdowns.MenuItemsProps{},
			dropdowns.DropdownMenuItem(
				dropdowns.MenuItemProps{},
				htmx.A(
					htmx.Attribute("href", "/me"),
					htmx.Text("Profile"),
				),
			),
			dropdowns.DropdownMenuItem(
				dropdowns.MenuItemProps{},
				htmx.A(
					htmx.Attribute("href", "/logout"),
					htmx.Text("Logout"),
				),
			),
		),
	)
}
