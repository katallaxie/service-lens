package partials

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/katallaxie/htmx/tables"
)

// ProfilePartialListControllerImpl ...
type ProfilePartialListControllerImpl struct {
	profiles tables.Results[models.Profile]
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.UnimplementedController
}

// NewProfilePartialListController ...
func NewProfilePartialListController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfilePartialListControllerImpl {
	return &ProfilePartialListControllerImpl{
		store: store,
	}
}

// Prepare ...
func (w *ProfilePartialListControllerImpl) Prepare() error {
	err := w.BindQuery(&w.profiles)
	if err != nil {
		return err
	}

	return w.store.ReadTx(w.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListProfiles(ctx, &w.profiles)
	})
}

// Get ...
func (w *ProfilePartialListControllerImpl) Get() error {
	return w.Render(
		htmx.Fragment(
			htmx.ForEach(w.profiles.GetRows(), func(e *models.Profile, profileIdx int) htmx.Node {
				return dropdowns.DropdownMenuItem(
					dropdowns.DropdownMenuItemProps{},
					htmx.A(
						htmx.Text(e.Name),
						htmx.DataAttribute("profile", e.ID.String()),
						htmx.HyperScript(`on click set (previous <input/>).value to my @data-profile then put my innerHTML into #profiles-button`),
					),
				)
			})...,
		),
	)
}
