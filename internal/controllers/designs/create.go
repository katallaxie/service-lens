package designs

import (
	"context"
	"fmt"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"
)

// CreateDesignControllerImpl ...
type CreateDesignControllerImpl struct {
	design models.Design
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// NewCreateDesignControllerImpl ...
func NewCreateDesignControllerImpl(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateDesignControllerImpl {
	return &CreateDesignControllerImpl{store: store}
}

// Clone ...
func (i *CreateDesignControllerImpl) Clone() handlers.Controller {
	return &CreateDesignControllerImpl{store: i.store}
}

// Prepare ...
func (l *CreateDesignControllerImpl) Prepare() error {
	err := l.BindAll(l.design)
	if err != nil {
		return err
	}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateDesign(ctx, &l.design)
	})
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (l *CreateDesignControllerImpl) Post() error {
	l.Redirect(fmt.Sprintf(utils.ShowDesigUrlFormat, l.design.ID))

	return nil
}
