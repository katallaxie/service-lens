package profiles

import (
	"context"

	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/profiles"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
)

// ProfileListControllerImpl ...
type ProfileListControllerImpl struct {
	profiles tables.Results[models.Profile]
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfilesListController ...
func NewProfilesListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileListControllerImpl {
	return &ProfileListControllerImpl{store: store}
}

// Prepare ...
func (w *ProfileListControllerImpl) Prepare() error {
	if err := w.BindQuery(&w.profiles); err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, &w.profiles)
	})
}

// Get ...
func (w *ProfileListControllerImpl) Get() error {
	return w.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        w.Path(),
				User:        w.Session().User,
				Development: w.IsDevelopment(),
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						profiles.ProfilesTable(
							profiles.ProfilesTableProps{
								Profiles: w.profiles.GetRows(),
								Offset:   w.profiles.GetOffset(),
								Limit:    w.profiles.GetLimit(),
								Total:    w.profiles.GetTotalRows(),
								URL:      w.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
