package handlers

import (
	"context"

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

// CreateTagParams represents the parameters for creating a new tag.
type CreateTagParams struct {
	Name  string `json:"name" form:"name" validate:"required,min=3,max=255"`
	Value string `json:"value" form:"value" validate:"required,min=3,max=255"`
}

// TagsHandler handles user-related routes.
type TagsHandler struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewTagsHandler returns a new TagsHandler.
func NewTagsHandler(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *TagsHandler {
	return &TagsHandler{store: store}
}

// CreateTag handles the tag creation page.
func (h *TagsHandler) CreateTag(c fiber.Ctx) (htmx.Node, error) {
	tag := models.Tag{}

	err := c.Bind().Body(&tag)
	if err != nil {
		return nil, err
	}

	err = h.store.ReadWriteTx(c.Context(), func(ctx context.Context, w ports.ReadWriteTx) error {
		return w.CreateTag(ctx, &tag)
	})
	if err != nil {
		return nil, err
	}

	return htmx.Text("created"), nil
}

// ListTags handles the tags listing page.
func (h *TagsHandler) ListTags(c fiber.Ctx) (htmx.Node, error) {
	s, err := goth.SessionFromContext(c)
	if err != nil {
		return nil, err
	}

	results := tables.Results[models.Tag]{
		SearchFields: []string{"Name"},
	}

	if err := c.Bind().Query(&results); err != nil {
		return nil, err
	}

	err = h.store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTags(ctx, &results)
	})
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
