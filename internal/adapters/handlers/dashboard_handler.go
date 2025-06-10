package handlers

import (
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/dashboard"
	"github.com/katallaxie/service-lens/internal/utils"

	"github.com/gofiber/fiber/v2"
	middleware "github.com/katallaxie/fiber-htmx"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/htmx/stats"
	"github.com/katallaxie/htmx/tailwind"
)

func Dashboard() middleware.CompFunc {
	return func(c *fiber.Ctx) (htmx.Node, error) {
		return components.DefaultLayout(
			components.DefaultLayoutProps{
				// Path:        d.Path(),
				// User:        d.Session().User,
				// Development: d.IsDevelopment(),
			},
			func() htmx.Node {
				return htmx.Fragment(
					dashboard.WelcomeCard(
						dashboard.WelcomeCardProps{},
					),
					stats.Stats(
						stats.Props{
							ClassNames: htmx.ClassNames{
								tailwind.Shadow: false,
								tailwind.M2:     true,
							},
						},
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Designs"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.HxGet(utils.DashboardStatsDesignUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Profiles"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.HxGet(utils.DashboardStatsProfileUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Workloads"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.HxGet(utils.DashboardStatsWorkloadUrlFormat),
								htmx.HxTrigger("load"),
								loading.Spinner(
									loading.SpinnerProps{},
								),
							),
						),
					),
				)
			},
		), nil
	}
}
