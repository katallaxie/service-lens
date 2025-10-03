package designs

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

// DesignTitleCardProps ...
type DesignTitleCardProps struct {
	ClassNames htmx.ClassNames
	Design     models.Design
	Markdown   string
}

// DesignTitleCard ...
func DesignTitleCard(props DesignTitleCardProps) htmx.Node {
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
		htmx.ID("title"),
		cards.Body(
			cards.BodyProps{},
			typography.H2(
				typography.Props{},
				htmx.Text(props.Design.Title),
			),
			cards.Actions(
				cards.ActionsProps{},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxGet(fmt.Sprintf(utils.EditTitleUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
