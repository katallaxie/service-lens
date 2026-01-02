package components

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/pkg/slices"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Children []htmx.Node
	Head     []htmx.Node
}

// Page is a whole document to output.
func Page(props PageProps, children ...htmx.Node) htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    props.Title,
			Language: "en",
			Head: slices.Append(
				props.Head,
				[]htmx.Node{
					htmx.Link(
						htmx.Href("https://cdn.jsdelivr.net/npm/daisyui@5"),
						htmx.Rel("stylesheet"),
						htmx.Type("text/css"),
					),
					htmx.Script(
						htmx.Src("https://unpkg.com/@htmx/htmx-dark-mode@latest/dist/index.js"),
						htmx.Type("module"),
					),
					htmx.Script(
						htmx.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
					),
					htmx.Imports(
						htmx.ImportsProp{
							Resolver: cache.New(jsdeliver.New()),
							Pkgs: []imports.ExactPackage{
								{
									Name:    "htmx.org",
									Version: "2.0.4",
								},
							},
							Requires: []imports.Require{
								{
									File: "dist/htmx.esm.js",
								},
							},
						},
					),
					htmx.Script(
						htmx.Type("module"),
						htmx.Raw(`import htmx from "htmx.org";`),
					),
				}...),
		},
		htmx.Body(
			htmx.Group(children...),
		),
	)
}
