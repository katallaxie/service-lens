package designs

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

var _ = htmx.Controller(&SearchTemplatesControllerImpl{})

// Search ...
type SearchTemplatesControllerImpl struct {
	Templates tables.Results[models.Template]
	store     seed.Database[ports.ReadTx, ports.ReadWriteTx]
	htmx.DefaultController
}

// NewSearchTemplatesController ...
func NewSearchTemplatesController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *SearchTemplatesControllerImpl {
	return &SearchTemplatesControllerImpl{
		Templates: tables.Results[models.Template]{SearchFields: []string{"name"}},
		store:     store,
	}
}

// Error ...
func (l *SearchTemplatesControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *SearchTemplatesControllerImpl) Prepare() error {
	var params struct {
		Template string `json:"template" form:"template" query:"template" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.Templates.Search = params.Template

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTemplates(ctx, &l.Templates)
	})
}

// Get ...
func (l *SearchTemplatesControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.Templates.GetRows(), func(e *models.Template, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(e.ID)),
					htmx.Text(e.Name),
				)
			})...,
		),
	)
}
