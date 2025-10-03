package templates

import (
	"fmt"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/htmx/typography"
)

// TemplateTitleCardProps ...
type TemplateTitleCardProps struct {
	ClassNames htmx.ClassNames
	Template   models.Template
	Markdown   string
}

// TemplateTitleCard ...
func TemplateTitleCard(props TemplateTitleCardProps) htmx.Node {
	return cards.CardBorder(
		cards.Props{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
			),
		},
		htmx.HxTarget("this"),
		htmx.HxSwap("outerHTML"),
		htmx.ID("name"),
		cards.Body(
			cards.BodyProps{},
			typography.H2(
				typography.Props{},
				htmx.Text(props.Template.Name),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditTemplateTitleUrlFormat, props.Template.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
