package designs

import (
	"context"
	"errors"

	"github.com/google/uuid"
	handlers "github.com/katallaxie/fiber-htmx/v3"
	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/designs"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/gorm"

	htmx "github.com/katallaxie/htmx"
)

// CreateDesignController ...
type CreateDesignController struct {
	store seed.Database[ports.ReadTx, ports.ReadWriteTx]
	handlers.UnimplementedController
}

// NewCreateDesignController ...
func NewCreateDesignController(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) *CreateDesignController {
	return &CreateDesignController{store: store}
}

// Clone ...
func (i *CreateDesignController) Clone() handlers.Controller {
	return &CreateDesignController{store: i.store}
}

// Get ...
func (l *CreateDesignController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        l.Path(),
				User:        l.Session().User,
				Development: l.IsDevelopment(),
				Head: []htmx.Node{
					htmx.Script(
						htmx.Src("https://unpkg.com/@github/markdown-toolbar-element@latest/dist/index.js"),
						htmx.Type("module"),
					),
				},
			},
			func() htmx.Node {
				params := struct {
					Template uuid.UUID `json:"template"`
				}{}
				errorx.Ignore(params, l.BindQuery(&params))

				template := models.Template{
					ID: params.Template,
				}

				err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTemplate(ctx, &template)
				})
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}

				return designs.DesignNewForm(
					designs.DesignNewFormProps{
						Template: template.Body,
					},
				)
			},
		),
	)
}
