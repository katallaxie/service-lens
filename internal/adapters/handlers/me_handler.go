package handlers

import (
	"github.com/gofiber/fiber/v2"
	middleware "github.com/katallaxie/fiber-htmx"
	reload "github.com/katallaxie/fiber-reload"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/service-lens/internal/components"
	goth "github.com/zeiss/fiber-goth"
)

// Me package handlers
func Me() middleware.CompFunc {
	return func(c *fiber.Ctx) (htmx.Node, error) {
		s, err := goth.SessionFromContext(c)
		if err != nil {
			return nil, err
		}

		return components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        c.Path(),
				User:        s.User,
				Development: reload.IsDevelopment(c.UserContext()),
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Profile"),
						),
						htmx.Form(
							htmx.HxPost("/me"),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Name"),
									),
								),

								forms.TextInputBordered(
									forms.TextInputProps{
										Name:     "username",
										Value:    s.User.Name,
										Disabled: true,
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("Your full nane as it will appear in the system."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Email"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:     "email",
										Value:    s.User.Email,
										Disabled: true,
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("Your email address. This is where we will send notifications."),
									),
								),
							),

							cards.Actions(
								cards.ActionsProps{},
								buttons.OutlinePrimary(
									buttons.ButtonProps{
										Disabled: true,
									},
									htmx.Attribute("type", "submit"),
									htmx.Text("Update Profile"),
								),
							),
						),
					),
				)
			},
		), nil
	}
}
