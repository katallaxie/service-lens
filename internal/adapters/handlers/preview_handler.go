package handlers

import (
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/service-lens/internal/builders"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type PreviewHandler struct{}

func NewPreviewHandler() *PreviewHandler {
	return &PreviewHandler{}
}

func (h *PreviewHandler) Preview(c *fiber.Ctx) (htmx.Node, error) {
	var form struct {
		Body string `json:"body" form:"body"`
	}

	err := c.BodyParser(&form)
	if err != nil {
		return nil, err
	}

	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
			renderer.WithNodeRenderers(util.Prioritized(builders.NewMarkdownBuilder(), 1)),
		),
		goldmark.WithExtensions(
			extension.GFM,
			emoji.Emoji,
		),
	)

	var b bytes.Buffer
	err = markdown.Convert([]byte(form.Body), &b)
	if err != nil {
		return nil, err
	}

	fmt.Println(b.String())

	return htmx.Raw(b.String()), nil
}
