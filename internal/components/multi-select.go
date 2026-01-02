package components

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/icons"
	"github.com/katallaxie/htmx/icons/heroicons"
)

// MultiSelectProps ...
type MultiSelectProps struct{}

// MultiSelect ...
func MultiSelect(props MultiSelectProps) htmx.Node {
	return htmx.Group(
		htmx.Input(htmx.Name("multi-select"), htmx.Type("checkbox")),
		forms.TextInputBordered(
			forms.TextInputProps{
				Name:        "search",
				Placeholder: "Begin Typing To Search Users...",
			},
			htmx.HxPost("/search"),
			htmx.HxTrigger("input changed delay:500ms, search"),
			htmx.HxTarget("#search-results"),
			htmx.HxIndicator(".htmx-indicator"),
		),
		dropdowns.Dropdown(
			dropdowns.Props{},
			dropdowns.Button(
				dropdowns.ButtonProps{
					ClassNames: htmx.ClassNames{
						"btn":             true,
						"btn-sm":          true,
						"btn-outline":     true,
						"w-full":          true,
						"justify-between": true,
					},
				},
				htmx.Text("Select Profile"),
				heroicons.ChevronUpDownDefaultOutline(icons.IconProps{}),

				//         <h3>
				//   Search Contacts
				//   <span class="htmx-indicator">
				//     <img src="/img/bars.svg"/> Searching...
				//    </span>
				// </h3>
				// <input class="form-control" type="search"
				//        name="search" placeholder="Begin Typing To Search Users..."
				//        hx-post="/search"
				//        hx-trigger="input changed delay:500ms, search"
				//        hx-target="#search-results"
				//        hx-indicator=".htmx-indicator">

				// <table class="table">
				//     <thead>
				//     <tr>
				//       <th>First Name</th>
				//       <th>Last Name</th>
				//       <th>Email</th>
				//     </tr>
				//     </thead>
				//     <tbody id="search-results">
				//     </tbody>
				// </table>
			),
		),
	)
}
