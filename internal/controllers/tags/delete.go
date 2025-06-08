package tags

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
)

// TagDeleteControllerImpl ...
type TagDeleteControllerImpl struct {
	ID    uuid.UUID `param:"id"`
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewTagDeleteController ...
func NewTagDeleteController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *TagDeleteControllerImpl {
	return &TagDeleteControllerImpl{store: store}
}

// Prepare ...
func (p *TagDeleteControllerImpl) Prepare() error {
	return p.BindParams(p)
}

// Delete ...
func (p *TagDeleteControllerImpl) Delete() error {
	return p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteTag(ctx, &models.Tag{ID: p.ID})
	})
}
