package workloads

import (
	"context"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
)

// WorkloadDeleteControllerImpl ...
type WorkloadDeleteControllerImpl struct {
	workload models.Workload
	store    seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkloadDeleteController ...
func NewWorkloadDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkloadDeleteControllerImpl {
	return &WorkloadDeleteControllerImpl{
		store: store,
	}
}

// Prepare ...
func (p *WorkloadDeleteControllerImpl) Prepare() error {
	err := p.BindParams(&p.workload)
	if err != nil {
		return err
	}

	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteWorkload(ctx, &p.workload)
	})
}

// Delete ...
func (p *WorkloadDeleteControllerImpl) Delete() error {
	return p.Redirect("/workloads")
}
