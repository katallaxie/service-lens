package me

import (
	handler "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/service-lens/internal/components"
)

// IndexController ...
type IndexController struct {
	handler.UnimplementedController
}

func (c *IndexController) Clone() handler.Controller {
	return &IndexController{}
}

// NewIndexController ...
func NewIndexController() *IndexController {
	return &IndexController{}
}

// Get ...
func (l *IndexController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				return cards.CardBorder(
					cards.Props{
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
										Value:    l.Session().User.Name,
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
										Value:    l.Session().User.Email,
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
		),
	)
}
