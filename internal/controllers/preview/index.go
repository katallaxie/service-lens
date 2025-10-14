package preview

import (
	"bytes"

	handler "github.com/katallaxie/fiber-htmx/v3"
	"github.com/katallaxie/htmx"
	"github.com/katallaxie/service-lens/internal/builders"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// IndexController ...
type IndexController struct {
	Body string `json:"body" form:"body"`
	handler.UnimplementedController
}

func (c *IndexController) Clone() handler.Controller {
	return &IndexController{}
}

// NewIndexController ...
func NewIndexController() *IndexController {
	return &IndexController{}
}

func (l *IndexController) Prepare() error {
	if err := l.BindAll(l); err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *IndexController) Post() error {
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
	err := markdown.Convert([]byte(l.Body), &b)
	if err != nil {
		return err
	}

	return l.Render(htmx.Raw(b.String()))
}
