package profiles

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"

	"github.com/go-playground/validator/v10"
	htmx "github.com/katallaxie/htmx"
)

const (
	listProfilesURL = "/profiles"
)

var validate *validator.Validate

// CreateProfileControllerImpl ...
type CreateProfileControllerImpl struct {
	profile models.Profile
	store   seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateProfileController ...
func NewCreateProfileController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateProfileControllerImpl {
	return &CreateProfileControllerImpl{store: store}
}

// Prepare ...
func (l *CreateProfileControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.BindBody(&l.profile)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.profile)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateProfile(ctx, &l.profile)
	})
}

// Post ...
func (l *CreateProfileControllerImpl) Post() error {
	return l.Redirect(listProfilesURL)
}
