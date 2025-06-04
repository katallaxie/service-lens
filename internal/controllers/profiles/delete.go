package profiles

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// ProfileDeleteControllerImpl ...
type ProfileDeleteControllerImpl struct {
	profile models.Profile
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewProfileDeleteController ...
func NewProfileDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *ProfileDeleteControllerImpl {
	return &ProfileDeleteControllerImpl{
		profile: models.Profile{},
		store:   store,
	}
}

// Prepare ...
func (p *ProfileDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.profile)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteProfile(ctx, &p.profile)
	})
}

// Delete ...
func (p *ProfileDeleteControllerImpl) Delete() error {
	return p.Redirect("/profiles")
}
