package designs

import (
	"fmt"

	"github.com/katallaxie/service-lens/internal/builders"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/conv"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"github.com/zeiss/fiber-goth/adapters"

	"go.abhg.dev/goldmark/mermaid"
)

// DesignBodyCardProps ...
type DesignBodyCardProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// Design ...
	Design models.Design
	// Markdown ...
	Markdown string
	// User ...
	User adapters.GothUser
}

// DesignBodyCard ...
func DesignBodyCard(props DesignBodyCardProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
			),
		},
		cards.Body(
			cards.BodyProps{},
			htmx.Div(
				htmx.ID("body"),
				htmx.Markdown(
					conv.Bytes(props.Design.Body),
					goldmark.WithRendererOptions(
						html.WithXHTML(),
						html.WithUnsafe(),
						renderer.WithNodeRenderers(
							util.Prioritized(
								builders.NewMarkdownBuilder(
									builders.WithTaskURL(fmt.Sprintf(utils.DesignTasksUrlFormat, props.Design.ID)),
								), 1),
						),
					),
					goldmark.WithExtensions(
						extension.GFM,
						emoji.Emoji,
						&mermaid.Extender{},
					),
				),
			),
			cards.Actions(
				cards.ActionsProps{
					ClassNames: htmx.ClassNames{
						tailwind.JustifyEnd:     false,
						tailwind.JustifyBetween: true,
					},
				},
				DesignReactions(
					DesignReactionsProps{
						User:   props.User,
						Design: props.Design,
					},
				),
				buttons.Button(
					buttons.ButtonProps{},
					htmx.HxSwap("outerHTML"),
					htmx.HxGet(fmt.Sprintf(utils.EditBodyUrlFormat, props.Design.ID)),
					htmx.Text("Edit"),
				),
			),
		),
	)
}
