package environments

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	htmx "github.com/zeiss/fiber-htmx"
	seed "github.com/zeiss/gorm-seed"
)

// EnvironmentDeleteControllerImpl ...
type EnvironmentDeleteControllerImpl struct {
	environment models.Environment
	store       seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewEnvironmentDeleteController ...
func NewEnvironmentDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *EnvironmentDeleteControllerImpl {
	return &EnvironmentDeleteControllerImpl{
		environment: models.Environment{},
		store:       store,
	}
}

// Prepare ...
func (p *EnvironmentDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.environment)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteEnvironment(ctx, &p.environment)
	})
}

// Delete ...
func (p *EnvironmentDeleteControllerImpl) Delete() error {
	return p.Redirect("/environments")
}
