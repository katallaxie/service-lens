package designs

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/modals"
	"github.com/katallaxie/service-lens/internal/utils"
)

// AddTagModalProps ...
type AddTagModalProps struct {
	DesignID uuid.UUID
}

// AddTagModal ...
func AddTagModal(props AddTagModalProps) htmx.Node {
	return modals.Modal(
		modals.Props{
			ID: "add_tag_modal",
		},
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.AddDesignTagUrlFormat, props.DesignID)),
			htmx.HxTrigger("submit"),
			htmx.HxOn("htmx:after-settle", "event.target.closest('dialog').close(), event.target.reset()"),
			htmx.HxSwap("none"),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextInputBordered(
					forms.TextInputProps{
						Name:        "name",
						Placeholder: "Provide tag name ...",
					},
					htmx.AutoComplete("off"),
				),
				forms.FormControlLabel(
					forms.FormControlLabelProps{},
					forms.FormControlLabelText(
						forms.FormControlLabelTextProps{
							ClassNames: htmx.ClassNames{
								"text-neutral-500": true,
							},
						},
						htmx.Text("Use a unique name to identify the tag that has between 3 and 255 characters."),
					),
				),
			),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextInputBordered(
					forms.TextInputProps{
						Name:        "value",
						Placeholder: "Provide a tag value ...",
					},
					htmx.AutoComplete("off"),
				),
				forms.FormControlLabel(
					forms.FormControlLabelProps{},
					forms.FormControlLabelText(
						forms.FormControlLabelTextProps{
							ClassNames: htmx.ClassNames{
								"text-neutral-500": true,
							},
						},
						htmx.Text("Use a unique value of the tag that has between 3 and 255 characters."),
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
					htmx.Text("Create"),
				),
			),
		),
	)
}
