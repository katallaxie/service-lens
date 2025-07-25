package workloads

import (
	"context"

	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/toasts"
	"github.com/katallaxie/pkg/conv"
	seed "github.com/zeiss/gorm-seed"
)

var _ = htmx.Controller(&SearchEnvironmentsControllerImpl{})

// Search ...
type SearchEnvironmentsControllerImpl struct {
	environments tables.Results[models.Environment]
	store        seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchEnvironmentsController ...
func NewSearchEnvironmentsController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchEnvironmentsControllerImpl {
	return &SearchEnvironmentsControllerImpl{
		environments: tables.Results[models.Environment]{SearchFields: []string{"name"}},
		store:        store,
	}
}

// Error ...
func (l *SearchEnvironmentsControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchEnvironmentsControllerImpl) Prepare() error {
	var params struct {
		Environment string `json:"environment" form:"environment" query:"environment" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.environments.Search = params.Environment

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListEnvironments(ctx, &l.environments)
	})
}

// Get ...
func (l *SearchEnvironmentsControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.environments.GetRows(), func(e *models.Environment, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(e.ID)),
					htmx.Text(e.Name),
				)
			})...,
		),
	)
}
