package designs

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"
)

// DesignTagProps ...
type DesignTagProps struct {
	ClassNames htmx.ClassNames
	DesignID   uuid.UUID
	Tag        models.Tag
}

// DesignTag ...
func DesignTag(props DesignTagProps) htmx.Node {
	return htmx.FormElement(
		htmx.ClassNames{
			tailwind.Flex:    true,
			tailwind.WFull:   true,
			tailwind.SpaceX4: true,
		},
		htmx.HxDelete(fmt.Sprintf(utils.RemoveDesignTagUrlFormat, props.DesignID, props.Tag.ID)),
		htmx.HxConfirm("Are you sure you want to remove this tag?"),
		htmx.HxDisabledElt("button"),
		forms.FormControl(
			forms.FormControlProps{
				ClassNames: htmx.ClassNames{},
			},
			forms.TextInputBordered(
				forms.TextInputProps{
					Value:    props.Tag.Name,
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
					htmx.Text("Key in a tag."),
				),
			),
		),
		forms.FormControl(
			forms.FormControlProps{
				ClassNames: htmx.ClassNames{},
			},
			forms.TextInputBordered(
				forms.TextInputProps{
					Value:    props.Tag.Value,
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
					htmx.Text("Value in a tag."),
				),
			),
		),
		buttons.Button(
			buttons.ButtonProps{
				Type: "submit",
			},
			icons.TrashOutline(
				icons.IconProps{},
			),
		),
	)
}
