package environments

import (
	"context"
	"fmt"

	"github.com/katallaxie/htmx/toasts"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	"github.com/go-playground/validator/v10"
	htmx "github.com/katallaxie/htmx"
)

var validate *validator.Validate

// CreateEnvironmentControllerImpl ...
type CreateEnvironmentControllerImpl struct {
	environment models.Environment
	store       seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateEnvironmentController ...
func NewCreateEnvironmentController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateEnvironmentControllerImpl {
	return &CreateEnvironmentControllerImpl{store: store}
}

// Error ...
func (l *CreateEnvironmentControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *CreateEnvironmentControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.environment)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.environment)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateEnvironment(ctx, &l.environment)
	})
}

// Post ...
func (l *CreateEnvironmentControllerImpl) Post() error {
	return l.Redirect(fmt.Sprintf("/environments/%s", l.environment.ID))
}
