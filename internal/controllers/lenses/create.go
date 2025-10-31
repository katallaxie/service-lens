package lenses

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"

	handlers "github.com/katallaxie/fiber-htmx/v3"
	seed "github.com/zeiss/gorm-seed"
)

// CreateController ...
type CreateController struct {
	lens  models.Lens
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// Clone ...
func (c *CreateController) Clone() handlers.Controller {
	return &CreateController{store: c.store}
}

// NewCreateController ...
func NewCreateController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateController {
	return &CreateController{store: store}
}

func (c *CreateController) Prepare() error {
	spec, err := c.Ctx().FormFile("spec")
	if err != nil {
		return err
	}
	file, err := spec.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	err = c.lens.UnmarshalJSON(buf.Bytes())
	if err != nil {
		return err
	}

	c.lens.IsDraft = true // first draft the

	err = c.store.ReadWriteTx(c.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateLens(ctx, &c.lens)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (c *CreateController) Post() error {
	c.Redirect(fmt.Sprintf(utils.ShowLensUrlFormat, c.lens.ID))
	return nil
}
