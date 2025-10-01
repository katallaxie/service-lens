package handlers

import (
	"context"

	goth "github.com/katallaxie/fiber-goth/v3"
	reload "github.com/katallaxie/fiber-reload/v3"
	"github.com/katallaxie/pkg/errorx"
	"github.com/katallaxie/service-lens/internal/components"
	"github.com/katallaxie/service-lens/internal/components/dashboard"
	"github.com/katallaxie/service-lens/internal/ports"
	"github.com/katallaxie/service-lens/internal/utils"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/pkg/conv"

	"github.com/gofiber/fiber/v3"
	middleware "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/htmx/stats"
	"github.com/katallaxie/htmx/tailwind"
)

// GetDashboard returns the dashboard index handler.
func GetDashboard() middleware.CompFunc {
	return func(c fiber.Ctx) (htmx.Node, error) {
		return components.DefaultLayout(
			components.DefaultLayoutProps{
				Path:        c.Path(),
				User:        errorx.Ignore(goth.SessionFromContext(c)).User,
				Development: reload.IsDevelopment(c),
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

// GetDashboardTotalDesigns returns the total number of designs in the dashboard.
func GetDashboardTotalDesigns(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) middleware.CompFunc {
	return func(c fiber.Ctx) (htmx.Node, error) {
		return htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfDesigns(ctx, &total)
				})
				errorx.Panic(err)

				return stats.Value(
					stats.ValueProps{},
					htmx.Text(conv.String(total)),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		), nil
	}
}

// GetDashboardTotalWorkloads returns the total number of workloads in the dashboard.
func GetDashboardTotalWorkloads(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) middleware.CompFunc {
	return func(c fiber.Ctx) (htmx.Node, error) {
		return htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfWorkloads(ctx, &total)
				})
				errorx.Panic(err)

				return stats.Value(
					stats.ValueProps{},
					htmx.Text(conv.String(total)),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		), nil
	}
}

// GetDashboardTotalProfiles returns the total number of profiles in the dashboard.
func GetDashboardTotalProfiles(store seed.Database[ports.ReadTx, ports.ReadWriteTx]) middleware.CompFunc {
	return func(c fiber.Ctx) (htmx.Node, error) {
		return htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := store.ReadTx(c.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfProfiles(ctx, &total)
				})
				errorx.Panic(err)

				return stats.Value(
					stats.ValueProps{},
					htmx.Text(conv.String(total)),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		), nil
	}
}
