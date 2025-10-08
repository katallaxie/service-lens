package tags

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	seed "github.com/zeiss/gorm-seed"
)

// DeleteTagController ...
type DeleteTagController struct {
	model models.Tag
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// NewDeleteTagController ...
func NewDeleteTagController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *DeleteTagController {
	return &DeleteTagController{store: store}
}

// Clone ...
func (i *DeleteTagController) Clone() handlers.Controller {
	return &DeleteTagController{store: i.store}
}

// Prepare ...
func (i *DeleteTagController) Prepare() error {
	if err := i.BindAll(&i.model); err != nil {
		return err
	}

	return i.store.ReadWriteTx(i.Context(), func(ctx context.Context, w ports.ReadWriteTx) error {
		return w.DeleteTag(ctx, &i.model)
	})
}

// Delete ...
func (p *DeleteTagController) Delete() error {
	return nil
}
