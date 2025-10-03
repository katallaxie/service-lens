package designs

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/service-lens/internal/models"
)

// DesignTagsCardProps ...
type DesignTagsCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignTagsCard ...
func DesignTagsCard(props DesignTagsCardProps) htmx.Node {
	return cards.CardBorder(
		cards.Props{
			ClassNames: htmx.ClassNames{
				tailwind.M2: true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Tags"),
			),
			htmx.Div(
				htmx.ID("tags"),
				htmx.Group(
					htmx.ForEach(props.Design.Tags, func(tag models.Tag, idx int) htmx.Node {
						return DesignTag(
							DesignTagProps{
								DesignID: props.Design.ID,
								Tag:      tag,
							},
						)
					},
					)...,
				),
			),
			AddTagModal(
				AddTagModalProps{
					DesignID: props.Design.ID,
				},
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{
						Type: "button",
					},
					htmx.OnClick("add_tag_modal.showModal()"),
					htmx.Text("Add Tag"),
				),
			),
		),
	)
}
