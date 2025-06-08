package designs

import (
	"fmt"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/avatars"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/htmx/tooltips"
	"github.com/katallaxie/htmx/typography"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/pkg/cast"
)

// DesignCommentProps ...
type DesignCommentProps struct {
	Comment models.DesignComment
	User    adapters.GothUser
	Design  models.Design
}

// DesignComment ...
func DesignComment(props DesignCommentProps) htmx.Node {
	return cards.CardBordered(
		cards.CardProps{
			ClassNames: htmx.ClassNames{
				tailwind.M2: true,
			},
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{
					ClassNames: htmx.ClassNames{
						tailwind.FontNormal: true,
						tailwind.TextBase:   true,
					},
				},
				tooltips.Tooltip(
					tooltips.TooltipProps{
						DataTip: props.User.Name,
					},
					avatars.AvatarRoundSmall(
						avatars.AvatarProps{},
						htmx.Img(
							htmx.Attribute("src", cast.Value(props.Comment.Author.Image)),
						),
					),
				),

				htmx.Text(fmt.Sprintf("commented on %s", props.Comment.CreatedAt.Format("Monday 02, 2006"))),
			),
			htmx.Text(props.Comment.Comment),
			cards.Actions(
				cards.ActionsProps{
					ClassNames: htmx.ClassNames{
						tailwind.JustifyEnd:     false,
						tailwind.JustifyBetween: true,
					},
				},
				DesignCommentReactions(
					DesignCommentReactionsProps{
						User:    props.User,
						Design:  props.Design,
						Comment: props.Comment,
					},
				),
				dropdowns.Dropdown(
					dropdowns.DropdownProps{
						ClassNames: htmx.ClassNames{},
					},
					dropdowns.DropdownButton(
						dropdowns.DropdownButtonProps{
							ClassNames: htmx.ClassNames{
								"btn": true,
							},
						},
						icons.EllipsisHorizontalOutline(
							icons.IconProps{},
						),
					),
					dropdowns.DropdownMenuItems(
						dropdowns.DropdownMenuItemsProps{
							ClassNames: htmx.ClassNames{
								tailwind.WFull: false,
								"w-52":         false,
							},
						},
						dropdowns.DropdownMenuItem(
							dropdowns.DropdownMenuItemProps{},
							htmx.A(
								typography.Error(
									typography.Props{},
									htmx.Text("Delete"),
								),
								htmx.HxDelete(fmt.Sprintf(utils.DeleteDesignCommentUrlFormat, props.Design.ID, props.Comment.ID)),
								htmx.HxTarget("closest .card"),
								htmx.HxSwap("outerHTML"),
								htmx.HxConfirm("Are you sure you want to delete this comment?"),
							),
						),
					),
				),
			),
		),
	)
}
