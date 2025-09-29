package dashboard

import (
	"github.com/katallaxie/fiber-goth/adapters"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/tailwind"
)

// WelcomeCardProps ...
type WelcomeCardProps struct {
	// ClassNames ...
	ClassNames htmx.ClassNames
	// User ...
	User adapters.GothUser
}

// WelcomeCard ...
func WelcomeCard(props WelcomeCardProps) htmx.Node {
	return cards.CardBorder(
		cards.CardProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					tailwind.M2: true,
				},
			),
		},
		cards.Body(
			cards.BodyProps{},
			cards.Title(
				cards.TitleProps{},
				htmx.Text("Welcome"),
			),
		),
	)
}
