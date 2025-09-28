package designs

import (
	"fmt"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/avatars"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/cast"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/utils"
	"github.com/zeiss/fiber-goth/adapters"
)

// DesignCommentsCardProps ...
type DesignCommentsCardProps struct {
	User       adapters.GothUser
	ClassNames htmx.ClassNames
	Design     models.Design
}

// DesignCommentsCard ...
func DesignCommentsCard(props DesignCommentsCardProps) htmx.Node {
	return htmx.Fragment(
		htmx.Div(
			htmx.ID("comments"),
			htmx.Group(htmx.ForEach(tables.RowsPtr(props.Design.Comments), func(c *models.DesignComment, choiceIdx int) htmx.Node {
				return DesignComment(
					DesignCommentProps{
						Comment: cast.Value(c),
						User:    props.User,
						Design:  props.Design,
					},
				)
			})...),
		),
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.CreateDesignCommentUrlFormat, props.Design.ID)),
			htmx.HxTarget("#comments"),
			htmx.HxSwap("beforeend"),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						avatars.AvatarRoundSmall(
							avatars.AvatarProps{},
							htmx.Img(
								htmx.Attribute("src", cast.Value(props.User.Image)),
							),
						),
						htmx.Text("Add a comment"),
					),
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						htmx.StyleElement(htmx.Raw(
							`.CodeMirror, .CodeMirror-scroll {
	min-height: 200px;
}`,
						)),
						htmx.Div(
							// 					alpine.XData(`{
							//     value: 'Start typing...',
							//     init() {
							//         let editor = new SimpleMDE({
							//           element: this.$refs.editor,
							// 		  status: false,
							//           previewRender: function(plainText, preview) {
							//             htmx.ajax('POST', '/preview', {values: {body: plainText}, target: '.editor-preview', swap: 'innerHTML'})

							//             return "Loading...";
							//           }
							//         })
							//         editor.value(this.value)
							//         editor.codemirror.on('change', () => {
							//             this.value = editor.value()
							//         })
							//     },
							// }`,
							// 					),
							forms.TextareaBordered(
								forms.TextareaProps{
									ClassNames: htmx.ClassNames{
										"h-[50vh]": true,
									},
									Name: "comment",
								},
								// alpine.XRef("editor"),
							),
						),
						forms.FormControlLabel(
							forms.FormControlLabelProps{},
							forms.FormControlLabelText(
								forms.FormControlLabelTextProps{
									ClassNames: htmx.ClassNames{
										"text-neutral-500": true,
									},
								},
								htmx.Text("Supports Markdown."),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Attribute("type", "submit"),
							htmx.Text("Comment"),
						),
					),
				),
			),
		),
	)
}
