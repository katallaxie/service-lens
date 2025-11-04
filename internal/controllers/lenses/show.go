package lenses

import (
	"context"
	"fmt"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/lenses"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"
)

// ShowControllerImpl ...
type ShowControllerImpl struct {
	lens  models.Lens
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// NewShowController ...
func NewShowController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ShowControllerImpl {
	return &ShowControllerImpl{
		store: store,
	}
}

// Clone ...
func (l *ShowControllerImpl) Clone() handlers.Controller {
	return &ShowControllerImpl{
		store: l.store,
	}
}

// Prepare ...
func (l *ShowControllerImpl) Prepare() error {
	err := l.BindAll(&l.lens)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetLens(ctx, &l.lens)
	})
}

// Get ...
func (l *ShowControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
			},
			func() htmx.Node {
				return htmx.Fragment(

					cards.CardBorder(
						cards.Props{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Name"),
								),
								htmx.H3(
									htmx.Text(l.lens.Name),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Version"),
								),
								htmx.H3(
									htmx.Text(conv.String(l.lens.Version)),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Description"),
								),
								htmx.H3(
									htmx.Text(l.lens.Description),
								),
							),
							lenses.LensesStatus(
								lenses.LensesStatusProps{
									IsDraft: l.lens.IsDraft,
								},
								htmx.ID("status"),
							),
							cards.Actions(
								cards.ActionsProps{},
								lenses.LensesPublishButton(
									lenses.LensesPublishButtonProps{
										ID:      l.lens.ID,
										IsDraft: l.lens.IsDraft,
									},
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(fmt.Sprintf(utils.DeleteLensUrlFormat, l.lens.ID)),
									htmx.HxConfirm("Are you sure you want to delete this lens?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
					lenses.LensMetadataCard(
						lenses.LensMetadataCardProps{
							Lens: l.lens,
						},
					),
				)
			},
		),
	)
}
