package workflows

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/toasts"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
)

// WorkflowDeleteControllerImpl ...
type WorkflowDeleteControllerImpl struct {
	ID    uuid.UUID `param:"id"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewWorkflowDeleteController ...
func NewWorkflowDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *WorkflowDeleteControllerImpl {
	return &WorkflowDeleteControllerImpl{store: store}
}

// Prepare ...
func (p *WorkflowDeleteControllerImpl) Prepare() error {
	return p.BindParams(p)
}

// Error ...
func (p *WorkflowDeleteControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Delete ...
func (p *WorkflowDeleteControllerImpl) Delete() error {
	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteWorkflow(ctx, &models.Workflow{ID: p.ID})
	})
}
