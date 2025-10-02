package designs

import (
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/htmx/modals"
	"github.com/katallaxie/htmx/tailwind"
)

// NewDesignModalProps ...
type NewDesignModalProps struct{}

// NewDesignModal ...
func NewDesignModal() htmx.Node {
	return modals.Modal(
		modals.Props{
			ID: "new_design_modal",
		},
		htmx.FormElement(
			htmx.ID("new-design-form"),
			htmx.Action(utils.CreateDesignUrlFormat),
			htmx.Method("get"),
			// htmx.HxDisabledElt("find button, find input"),
			// htmx.HxOn("htmx:after-settle", "event.target.closest('dialog').close(), event.target.reset()"),
			forms.FormControl(
				forms.FormControlProps{},
				htmx.Div(
					htmx.ClassNames{
						tailwind.Flex:           true,
						tailwind.JustifyBetween: true,
					},
					forms.Datalist(
						forms.DatalistProps{
							ID:          "templates",
							Name:        "template",
							Placeholder: "Search a template ...",
							URL:         utils.DesignSearchTemplatesUrlFormat,
						},
					),
					loading.Spinner(
						loading.SpinnerProps{
							ClassNames: htmx.ClassNames{
								"htmx-indicator": true,
								tailwind.M2:      true,
							},
						},
					),
				),
				forms.FormControlLabel(
					forms.FormControlLabelProps{},
					forms.FormControlLabelText(
						forms.FormControlLabelTextProps{
							ClassNames: htmx.ClassNames{
								"text-neutral-500": true,
							},
						},
						htmx.Text("Optional - The template to use for the new design"),
					),
				),
			),
			modals.ModalAction(
				modals.ModalActionProps{},
				buttons.Ghost(
					buttons.ButtonProps{
						Type: "button",
					},
					htmx.Text("Cancel"),
					htmx.Attribute("formnovalidate", ""),
					htmx.OnClick("event.target.closest('dialog').close()"),
				),
				buttons.Button(
					buttons.ButtonProps{
						Type: "submit",
					},
					htmx.Text("Add Design"),
				),
			),
		),
	)
}
