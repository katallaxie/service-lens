package designs_edit_body

import (
	"context"
	"fmt"

	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/service-lens/internal/components/designs"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
)

var _ = htmx.Controller(&ShowControllerImpl{})

// ShowControllerImpl ...
type ShowControllerImpl struct {
	Design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEditController ...
func NewEditController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowControllerImpl {
	return &ShowControllerImpl{store: store}
}

// Prepare ...
func (l *ShowControllerImpl) Prepare() error {
	err := l.BindParams(&l.Design)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetDesign(ctx, &l.Design)
	})
}

// Prepare ...
func (l *ShowControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.ID("body"),
				htmx.HxSwapOob(conv.String(true)),
				htmx.FormElement(
					forms.FormControl(
						forms.FormControlProps{
							ClassNames: htmx.ClassNames{},
						},
						designs.Editor(
							designs.EditorProps{
								Content: l.Design.Body,
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
								htmx.Text("Supports Markdown."),
							),
						),
					),
				),
			),
			buttons.Button(
				buttons.ButtonProps{
					Type: "submit",
				},
				htmx.HxSwap("outerHTML"),
				htmx.HxPut(fmt.Sprintf(utils.EditBodyUrlFormat, l.Design.ID)),
				htmx.HxInclude("body"),
				htmx.Text("Update"),
			),
		),
	)
}
