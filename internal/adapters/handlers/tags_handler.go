package handlers

import (
	"github.com/gofiber/fiber/v3"
	goth "github.com/katallaxie/fiber-goth/v3"
	reload "github.com/katallaxie/fiber-reload/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/htmx/tailwind"
	"github.com/katallaxie/pkg/dbx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/tags"
	"github.com/katallaxie/service-lens/internal/models"
	"github.com/katallaxie/service-lens/internal/ports"
)

// TagsHandler handles user-related routes.
type TagsHandler struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewTagsHandler returns a new TagsHandler.
func NewTagsHandler(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *TagsHandler {
	return &TagsHandler{store: store}
}

// ListTags handles the tags listing page.
func (h *TagsHandler) ListTags(c fiber.Ctx) (htmx.Node, error) {
	s, err := goth.SessionFromContext(c)
	if err != nil {
		return nil, err
	}

	return components.DefaultLayout(
		components.DefaultLayoutProps{
			Path:        c.Path(),
			User:        s.User,
			Development: reload.IsDevelopment(c),
		},
		func() htmx.Node {
			results := tables.Results[models.Tag]{
				SearchFields: []string{"Name"},
			}

			// errorx.Panic(c.BindQuery(&results))
			// errorx.Panic(c.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
			// 	return tx.ListTags(ctx, &results)
			// }))

			return cards.CardBorder(
				cards.Props{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					tags.TagsTable(
						tags.TagsTableProps{
							Tags:   results.GetRows(),
							Offset: results.GetOffset(),
							Limit:  results.GetLimit(),
							Total:  results.GetTotalRows(),
							URL:    c.OriginalURL(),
						},
					),
				),
			)
		},
	), nil
}
