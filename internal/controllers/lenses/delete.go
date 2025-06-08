package lenses

import (
	"context"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
)

// LensDeleteControllerImpl ...
type LensDeleteControllerImpl struct {
	lens  models.Lens
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewLensDeleteController ...
func NewLensDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *LensDeleteControllerImpl {
	return &LensDeleteControllerImpl{
		store: store,
	}
}

// Prepare ...
func (p *LensDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.lens)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteLens(ctx, &p.lens)
	})
}

// Delete ...
func (p *LensDeleteControllerImpl) Delete() error {
	return p.Redirect("/lenses")
}
