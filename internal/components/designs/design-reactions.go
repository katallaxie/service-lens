package designs

import (
	"fmt"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"
	"github.com/zeiss/fiber-goth/adapters"
)

// DesignReactionsProps ...
type DesignReactionsProps struct {
	User   adapters.GothUser
	Design models.Design
}

// DesignReactions ...
func DesignReactions(props DesignReactionsProps) htmx.Node {
	return htmx.Div(
		htmx.ID("design-reactions"),
		htmx.HxSwapOob(conv.String(true)),
		htmx.ClassNames{
			tailwind.Flex:        true,
			tailwind.ItemsCenter: true,
		},
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.CreateDesignReactionUrlFormat, props.Design.ID)),
			components.EmojiPicker(
				components.EmojiPickerProps{},
			),
		),
		htmx.Group(
			htmx.Map(props.Design.GetReactionsByValue(), func(reaction string, reactions []models.Reaction) htmx.Node {
				react := slices.Index(func(reaction models.Reaction) bool {
					return reaction.ReactorID == props.User.ID
				}, reactions...)
				return htmx.FormElement(
					htmx.IfElse(
						react != -1,
						htmx.HxDelete(fmt.Sprintf(utils.DeleteDesignReactionUrlFormat, props.Design.ID, reactions[react].ID)),
						htmx.HxPost(fmt.Sprintf(utils.CreateDesignReactionUrlFormat, props.Design.ID)),
					),
					htmx.Input(
						htmx.Type("hidden"),
						htmx.Name("reaction"),
						htmx.Value(reaction),
					),
					buttons.Button(
						buttons.ButtonProps{
							ClassNames: htmx.ClassNames{
								tailwind.Mx1: true,
							},
						},
						htmx.Text(fmt.Sprintf("%s (%d)", reaction, (len(reactions)))),
					),
				)
			},
			)...,
		),
	)
}
