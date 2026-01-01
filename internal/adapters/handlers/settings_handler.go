package handlers

import (
	"github.com/gofiber/fiber/v3"
	goth "github.com/katallaxie/fiber-goth/v3"
	reload "github.com/katallaxie/fiber-reload/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/collapsible"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/joins"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

func (h *SettingsHandler) ListSettings(c fiber.Ctx) (htmx.Node, error) {
	return components.DefaultLayout(
		components.DefaultLayoutProps{
			Path:        c.Path(),
			User:        errorx.Ignore(goth.SessionFromContext(c)).User,
			Development: reload.IsDevelopment(c),
		},
		func() htmx.Node {
			return htmx.Fragment(
				cards.CardBorder(
					cards.Props{
						ClassNames: htmx.Merge(
							htmx.ClassNames{
								tailwind.M2: true,
							},
						),
					},
					htmx.Form(
						htmx.HxPut(""),
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Settings"),
							),
							joins.Vertical(
								joins.Props{
									ClassNames: htmx.ClassNames{
										tailwind.WFull: true,
									},
								},
								collapsible.Collapse(
									collapsible.Props{
										ClassNames: htmx.ClassNames{
											"join-item": true,
										},
									},
									collapsible.Checkbox(
										collapsible.CheckboxProps{},
									),
									collapsible.Title(
										collapsible.TitleProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Microsoft Entra ID"),
									),
									collapsible.Content(
										collapsible.ContentProps{},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{
													"flex":            true,
													"justify-between": true,
													"flex-row":        true,
												},
											},
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{},
													htmx.Text("Enable"),
												),
											),
											forms.Toggle(
												forms.ToggleProps{
													Name:  "entra_id_enabled",
													Value: "true",
												},
											),
										),
									),
								),
								collapsible.Collapse(
									collapsible.Props{
										ClassNames: htmx.ClassNames{
											"join-item": true,
										},
									},
									collapsible.Checkbox(
										collapsible.CheckboxProps{},
									),
									collapsible.Title(
										collapsible.TitleProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("GitHub"),
									),
									collapsible.Content(
										collapsible.ContentProps{},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{
													"flex":            true,
													"justify-between": true,
													"flex-row":        true,
												},
											},
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{},
													htmx.Text("Enable"),
												),
											),
											forms.Toggle(
												forms.ToggleProps{
													Name:  "github_enabled",
													Value: "true",
												},
											),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Save"),
								),
							),
						),
					),
				),
			)
		},
	), nil
}
