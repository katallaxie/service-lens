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

var _ = htmx.Controller(&SearchLensesControllerImpl{})

// Search ...
type SearchLensesControllerImpl struct {
	lenses tables.Results[models.Lens]
	store  seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchLensesController ...
func NewSearchLensesController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchLensesControllerImpl {
	return &SearchLensesControllerImpl{
		lenses: tables.Results[models.Lens]{SearchFields: []string{"name"}},
		store:  store,
	}
}

// Error ...
func (l *SearchLensesControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchLensesControllerImpl) Prepare() error {
	var params struct {
		Lens string `json:"lens" form:"lens" query:"lens" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.lenses.Search = params.Lens

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListLenses(ctx, &l.lenses)
	})
}

// Get ...
func (l *SearchLensesControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.lenses.GetRows(), func(e *models.Lens, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(e.ID)),
					htmx.Text(e.Name),
				)
			})...,
		),
	)
}
