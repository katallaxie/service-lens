package handlers

import (
	"github.com/gofiber/fiber/v3"
	htmx "github.com/katallaxie/htmx"
)

type PreviewHandler struct{}

func NewPreviewHandler() *PreviewHandler {
	return &PreviewHandler{}
}

func (h *PreviewHandler) Preview(c fiber.Ctx) (htmx.Node, error) {
	// var form struct {
	// 	Body string `json:"body" form:"body"`
	// }

	// err := c.BodyParser(&form)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil

	// markdown := goldmark.New(
	// 	goldmark.WithRendererOptions(
	// 		html.WithXHTML(),
	// 		html.WithUnsafe(),
	// 		renderer.WithNodeRenderers(util.Prioritized(builders.NewMarkdownBuilder(), 1)),
	// 	),
	// 	goldmark.WithExtensions(
	// 		extension.GFM,
	// 		emoji.Emoji,
	// 	),
	// )

	// var b bytes.Buffer
	// err = markdown.Convert([]byte(form.Body), &b)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(b.String())

	// return htmx.Raw(b.String()), nil
}
