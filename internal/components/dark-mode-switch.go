package components

import (
	htmx "github.com/katallaxie/htmx"
	darkmode "github.com/katallaxie/htmx-dark-mode-element"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/icons/heroicons"
	"github.com/katallaxie/htmx/swaps"
)

// DarkModeSwitchProps ...
type DarkModeSwitchProps struct {
	htmx.ClassNames
}

// DarkModeSwitch ...
func DarkModeSwitch(props DarkModeSwitchProps) htmx.Node {
	return darkmode.DarkMode(
		darkmode.DarkTheme("dark"),
		darkmode.LightTheme("light"),
		buttons.Circle(
			buttons.ButtonProps{},
			swaps.Swap(
				swaps.Props{},
				htmx.Input(
					htmx.Attribute("type", "checkbox"),
				),
				swaps.On(
					swaps.Props{},
					heroicons.MoonDefaultOutline(
						icons.IconProps{},
					),
				),
				swaps.Off(
					swaps.Props{},
					heroicons.SunDefaultOutline(
						icons.IconProps{},
					),
				),
			),
		),
	)
}
