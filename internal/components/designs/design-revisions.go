package designs

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/collapsible"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/service-lens/internal/utils"
)

// DesignRevisionsCardProps ...
type DesignRevisionsCardProps struct {
	ClassNames htmx.ClassNames
	DesignID   uuid.UUID
}

// DesignRevisionsCard ...
func DesignRevisionsCard(props DesignRevisionsCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
				props.ClassNames,
			),
		},
		cards.Body(
			cards.BodyProps{},
			collapsible.CollapseArrow(
				collapsible.CollapseProps{},
				htmx.HxTrigger("click once"),
				htmx.HxGet(fmt.Sprintf(utils.ListDesignRevisionsUrlFormat, props.DesignID)),
				htmx.HxTarget(".collapse-content"),
				htmx.HxIndicator("find .htmx-indicator"),
				collapsible.CollapseCheckbox(
					collapsible.CollapseCheckboxProps{},
				),
				collapsible.CollapseTitle(
					collapsible.CollapseTitleProps{
						ClassNames: htmx.ClassNames{
							tailwind.Flex:        true,
							tailwind.ItemsCenter: true,
						},
					},
					htmx.Text("Revisions"),
					loading.SpinnerSmall(
						loading.SpinnerProps{
							ClassNames: htmx.ClassNames{
								tailwind.Mx2:     true,
								"htmx-indicator": true,
							},
						},
					),
				),
				collapsible.CollapseContent(
					collapsible.CollapseContentProps{},
					htmx.Div(
						htmx.Text("Loading..."),
					),
				),
			),
		),
	)
}
