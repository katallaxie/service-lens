package workflows

import (
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/modals"
)

// NewWorkflowModalProps ...
type NewWorkflowModalProps struct{}

// NewWorkflowModal ...
func NewWorkflowModal() htmx.Node {
	return modals.Modal(
		modals.ModalProps{
			ID: "new_workflow_modal",
		},
		htmx.FormElement(
			htmx.HxPost(utils.CreateWorkflowUrlFormat),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextInputBordered(
					forms.TextInputProps{
						Name:        "name",
						Placeholder: "Provide workflow name ...",
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
						htmx.Text("Use a unique value of the workflow name that has between 3 and 255 characters."),
					),
				),
			),
			forms.FormControl(
				forms.FormControlProps{
					ClassNames: htmx.ClassNames{},
				},
				forms.TextareaBordered(
					forms.TextareaProps{
						Name:        "description",
						Placeholder: "Provider a workflow description ...",
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
						htmx.Text("Provide a description of the workflow that has between 3 and 1024 characters."),
					),
				),
			),
			modals.ModalAction(
				modals.ModalActionProps{},
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
