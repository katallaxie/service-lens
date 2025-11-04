package tags

import (
	"context"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"
)

// CreateController ...
type CreateController struct {
	model models.Tag
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// Clone ...
func (i *CreateController) Clone() handlers.Controller {
	return &CreateController{store: i.store}
}

// NewCreateController ...
func NewCreateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateController {
	return &CreateController{store: store}
}

// Prepare ...
func (i *CreateController) Prepare() error {
	if err := i.BindBody(&i.model); err != nil {
		return err
	}

	return i.store.ReadWriteTx(i.Context(), func(ctx context.Context, w ports.ReadWriteTx) error {
		return w.CreateTag(ctx, &i.model)
	})
}

// Post ...
func (i *CreateController) Post() error {
	i.Redirect(utils.ListTagsUrlFormat)

	return nil
}
