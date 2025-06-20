package designs

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/katallaxie/htmx/toasts"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"

	htmx "github.com/katallaxie/htmx"
	seed "github.com/zeiss/gorm-seed"
)

var validate *validate.Validator

type CreateDesignBody struct {
	Title      string    `json:"title" form:"title" validate:"required,min=3,max=2048"`
	Body       string    `json:"body" form:"body" validate:"required"`
	WorkflowID uuid.UUID `json:"workflow_id" form:"omitempty,workflow_id" validate:"omitempty,required,uuid"`
	Tags       []struct {
		Name  string `json:"name" form:"name" validate:"required"`
		Value string `json:"value" form:"value" validate:"required"`
	} `json:"tags" form:"tags"`
}

// CreateDesignControllerImpl ...
type CreateDesignControllerImpl struct {
	body  CreateDesignBody
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewCreateDesignController ...
func NewCreateDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateDesignControllerImpl {
	return &CreateDesignControllerImpl{store: store}
}

// Prepare ...
func (l *CreateDesignControllerImpl) Prepare() error {
	validate = validate.New()

	err := l.BindBody(&l.body)
	if err != nil {
		return err
	}

	err = validate.Struct(&l.body)
	if err != nil {
		return err
	}

	return nil
}

// Error ...
func (l *CreateDesignControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
func (l *CreateDesignControllerImpl) Post() error {
	design := models.Design{
		Title:    l.body.Title,
		Body:     l.body.Body,
		AuthorID: l.Session().UserID,
	}

	if l.body.WorkflowID != uuid.Nil {
		design.Workable = &models.Workable{
			WorkableType: models.WorkableTypeDesign,
			WorkflowTransition: models.WorkflowTransition{
				WorkflowID: l.body.WorkflowID,
			},
		}
	}

	for _, tag := range l.body.Tags {
		design.Tags = append(design.Tags, models.Tag{
			Name:  tag.Name,
			Value: tag.Value,
		})
	}

	err := l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateDesign(ctx, &design)
	})
	if err != nil {
		return err
	}

	return l.Redirect(fmt.Sprintf(utils.ShowDesigUrlFormat, design.ID))
}
