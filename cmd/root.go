package cmd

import (
	"context"
	"fmt"

	"github.com/katallaxie/service-lens/internal/adapters/db"
	"github.com/katallaxie/service-lens/internal/adapters/handlers"
	config "github.com/katallaxie/service-lens/internal/cfg"
	"github.com/katallaxie/service-lens/internal/controllers/designs"
	"github.com/katallaxie/service-lens/internal/controllers/environments"
	"github.com/katallaxie/service-lens/internal/controllers/lenses"
	"github.com/katallaxie/service-lens/internal/controllers/login"
	"github.com/katallaxie/service-lens/internal/controllers/me"
	"github.com/katallaxie/service-lens/internal/controllers/preview"
	"github.com/katallaxie/service-lens/internal/controllers/profiles"
	"github.com/katallaxie/service-lens/internal/controllers/tags"
	"github.com/katallaxie/service-lens/internal/controllers/workflows"
	"github.com/katallaxie/service-lens/internal/controllers/workloads"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	logger "github.com/gofiber/fiber/v3/middleware/logger"
	requestid "github.com/gofiber/fiber/v3/middleware/requestid"
	goth "github.com/katallaxie/fiber-goth/v3"
	adapter "github.com/katallaxie/fiber-goth/v3/adapters/gorm"
	"github.com/katallaxie/fiber-goth/v3/providers"
	"github.com/katallaxie/fiber-goth/v3/providers/entraid"
	"github.com/katallaxie/fiber-goth/v3/providers/github"
	htmx "github.com/katallaxie/fiber-htmx/v3"
	reload "github.com/katallaxie/fiber-reload/v3"
	"github.com/katallaxie/pkg/dbx"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	os      = "unknown"
	arch    = "unknown"
)

var versionFmt = fmt.Sprintf("%s-%s (%s) %s/%s", version, commit, date, os, arch)

var cfg = config.New()

func Init() error {
	ctx := context.Background()

	err := cfg.InitDefaultConfig()
	if err != nil {
		return err
	}

	Root.AddCommand(Migrate)
	Root.SilenceUsage = true

	err = Root.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

var Root = &cobra.Command{
	Use:     "service-lens",
	Short:   "Service Lens is a tool to help you manage your services.",
	Long:    `Service Lens is a tool to help you manage your services.`,
	Version: versionFmt,
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewWebSrv(cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *config.Config
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *config.Config) *WebSrv {
	return &WebSrv{cfg}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		if s.cfg.Flags.GitHubEnabled {
			providers.RegisterProvider(github.New(s.cfg.Flags.GitHubClientID, s.cfg.Flags.GitHubClientSecret, s.cfg.Flags.GitHubCallbackURL))
		}

		if s.cfg.Flags.EntraIDEnabled {
			providers.RegisterProvider(entraid.New(s.cfg.Flags.EntraIDClientID, s.cfg.Flags.EntraIDClientSecret, s.cfg.Flags.EntraIDCallbackURL, entraid.TenantType(s.cfg.Flags.EntraIDTenantID)))
		}

		conn, err := gorm.Open(postgres.Open(s.cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: s.cfg.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		store, err := dbx.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
		if err != nil {
			return err
		}

		gorm := adapter.New(conn)

		gothConfig := goth.Config{
			Adapter:        gorm,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
		}

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(helmet.New())
		app.Use(logger.New())
		app.Use(reload.Environment(s.cfg.Flags.Environment))

		if s.cfg.Flags.Environment == reload.Development {
			reload.WithHotReload(app)
		}

		app.Use(goth.NewProtectMiddleware(gothConfig))

		compFuncConfig := htmx.Config{
			ErrorHandler: htmx.ToastsErrorHandler,
		}

		app.Get("/", htmx.NewCompFuncHandler(handlers.GetDashboard(), compFuncConfig))
		app.Get("/login", htmx.NewControllerHandler(login.NewIndexLoginController(), compFuncConfig))
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))

		// User
		app.Get("/me", htmx.NewControllerHandler(me.NewIndexController(), compFuncConfig))

		// Preview
		app.Post("/preview", htmx.NewControllerHandler(preview.NewIndexController(), compFuncConfig))

		// Stats ...
		stats := app.Group("/stats")
		stats.Get("/profiles", htmx.NewCompFuncHandler(handlers.GetDashboardTotalProfiles(store), compFuncConfig))
		stats.Get("/workloads", htmx.NewCompFuncHandler(handlers.GetDashboardTotalWorkloads(store), compFuncConfig))
		stats.Get("/designs", htmx.NewCompFuncHandler(handlers.GetDashboardTotalDesigns(store), compFuncConfig))

		// Tags ...
		tg := app.Group("/tags")
		tg.Get("/", htmx.NewControllerHandler(tags.NewIndexController(store), compFuncConfig))
		tg.Post("/new", htmx.NewControllerHandler(tags.NewCreateController(store), compFuncConfig))
		tg.Delete("/:id", htmx.NewControllerHandler(tags.NewDeleteTagController(store), compFuncConfig))

		// Designs ...
		dg := app.Group("/designs")
		dg.Get("/", htmx.NewControllerHandler(designs.NewIndexController(store), compFuncConfig))
		dg.Get("/new", htmx.NewControllerHandler(designs.NewDesignController(store), compFuncConfig))
		dg.Post("/new", htmx.NewControllerHandler(designs.NewCreateDesignControllerImpl(store), compFuncConfig))
		dg.Get("/:id", htmx.NewControllerHandler(designs.NewShowDesignController(store), compFuncConfig))

		// designs.Post("/new", handlers.CreateDesign())
		// designs.Get("/search/workflows", handlers.SearchWorkflows())
		// designs.Get("/search/templates", handlers.SearchTemplates())
		// designs.Get("/:id", handlers.ShowDesign())
		// designs.Put("/:id", handlers.UpdateDesign())
		// designs.Delete("/:id", handlers.DeleteDesign())
		// designs.Post("/:id/tags", handlers.AddTagDesign())
		// designs.Delete("/:id/tags/:tag_id", handlers.RemoveTagDesign())
		// designs.Post("/:id/comments", handlers.CreateDesignComment())
		// designs.Delete("/:id/comments/:comment_id", handlers.DeleteDesignComment())
		// designs.Get("/:id/revisions", handlers.ListDesignRevisions())
		// designs.Get("/:id/body/edit", handlers.EditBodyDesign())
		// designs.Put("/:id/body/edit", handlers.UpdateBodyDesign())
		// designs.Get("/:id/title/edit", handlers.EditTitleDesign())
		// designs.Put("/:id/title/edit", handlers.UpdateTitleDesign())
		// designs.Post("/:id/reactions", handlers.DesignReactions())
		// designs.Post("/:id/tasks", handlers.Task())
		// designs.Delete("/:id/reactions/:reaction_id", handlers.DesignReactions())
		// designs.Post("/:id/comments/:comment_id/reactions", handlers.CreateDesignCommentReaction())
		// designs.Delete("/:id/comments/:comment_id/reactions/:reaction_id", handlers.DeleteDesignCommentReaction())

		// Profiles
		pg := app.Group("/profiles")
		pg.Get("/", htmx.NewControllerHandler(profiles.NewListController(store), compFuncConfig))
		// profiles.Get("/new", handlers.NewProfile())
		// profiles.Post("/new", handlers.CreateProfile())
		// profiles.Get("/:id", handlers.ShowProfile())
		// profiles.Put("/:id", handlers.EditProfile())
		// profiles.Delete("/:id", handlers.DeleteProfile())

		// Environments ...
		eg := app.Group("/environments")
		eg.Get("/", htmx.NewControllerHandler(environments.NewListController(store), compFuncConfig))
		// environments.Get("/new", handlers.NewEnvironment())
		// environments.Post("/new", handlers.CreateEnvironment())
		// environments.Get("/:id", handlers.ShowEnvironment())
		// environments.Get("/:id/edit", handlers.EditEnvironment())
		// environments.Put("/:id", handlers.UpdateEnvironment())
		// environments.Delete("/:id", handlers.DeleteEnvironment())

		// // Lenses ...
		lg := app.Group("/lenses")
		lg.Get("/", htmx.NewControllerHandler(lenses.NewListController(store), compFuncConfig))
		// lenses.Post("/", handlers.NewLens())
		// lenses.Get("/:id", handlers.ShowLens())
		// lenses.Get("/:id/edit", handlers.EditLens())
		// lenses.Put("/:id", handlers.UpdateLens())
		// lenses.Delete("/:id", handlers.DeleteLens())
		// lenses.Post("/:id/publish", handlers.PublishLens())
		// lenses.Delete("/:id/publish", handlers.UnpublishLens())

		// Workloads ...
		wg := app.Group("/workloads")
		wg.Get("/", htmx.NewControllerHandler(workloads.NewIndexController(store), compFuncConfig))
		// workloads.Get("/new", handlers.NewWorkload())
		// workloads.Post("/new", handlers.CreateWorkload())
		// workloads.Get("/search/lenses", handlers.SearchLenses())
		// workloads.Get("/search/environments", handlers.SearchEnvironments())
		// workloads.Get("/search/profiles", handlers.SearchProfiles())
		// workloads.Get("/:id", handlers.ShowWorkload())
		// workloads.Get("/:id/edit", handlers.EditWorkload())
		// workloads.Post("/:id/edit", handlers.EditWorkload())
		// workloads.Post("/:id/tags", handlers.AddTagWorkload())
		// workloads.Delete("/:id/tags/:tag_id", handlers.RemoveTagWorkload())
		// // app.Put("/workloads/:id", handlers.UpdateWorkload())
		// workloads.Delete("/:id", handlers.DeleteWorkload())
		// workloads.Get("/partials/environments", handlers.ListEnvironmentsPartial())
		// workloads.Get("/partials/profiles", handlers.ListProfilesPartial())
		// workloads.Get("/:id/lenses/:lens", handlers.ShowWorkloadLens())
		// workloads.Get("/:id/lenses/:lens/edit", handlers.EditWorkloadLens())
		// workloads.Get("/:workload/lenses/:lens/question/:question", handlers.ShowLensQuestion())
		// workloads.Put("/:workload/lenses/:lens/question/:question", handlers.UpdateWorkloadAnswer())

		// Workflows ...
		wfg := app.Group("/workflows")
		wfg.Get("/", htmx.NewControllerHandler(workflows.NewListController(store), compFuncConfig))
		// workflows.Post("/new", handlers.CreateWorkflow())
		// workflows.Get("/:id", handlers.ShowWorkflow())
		// workflows.Post("/:id/steps", handlers.CreateWorkflowStep())
		// workflows.Delete("/:id/steps/:step_id", handlers.DeleteWorkflowStep())
		// workflows.Put("/:id/steps", handlers.UpdateWorkflowSteps())
		// workflows.Delete("/:id", handlers.DeleteWorkflow())

		// // Templates ...
		// templates := app.Group("/templates")
		// templates.Get("/", handlers.ListTemplates())
		// templates.Get("/new", handlers.NewTemplate())
		// templates.Get("/:id", handlers.ShowTemplate())
		// templates.Delete("/:id", handlers.DeleteTemplate())
		// templates.Get("/:id/edit/body", handlers.EditTemplateBody())
		// templates.Put("/:id/edit/body", handlers.EditTemplateBody())
		// templates.Get("/:id/edit/title", handlers.EditTemplateTitle())
		// templates.Put("/:id/edit/title", handlers.EditTemplateTitle())
		// templates.Post("/new", handlers.CreateTemplate())

		// // Settings ...
		// app.Get("/settings", htmx.NewCompFuncHandler(settingsHandlers.ListSettings, compFuncConfig))

		// // Preview ...
		// app.Post("/preview", htmx.NewCompFuncHandler(previewHandlers.Preview, compFuncConfig))

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
