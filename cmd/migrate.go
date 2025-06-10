package cmd

import (
	"github.com/katallaxie/service-lens/internal/adapters/db"
	"github.com/katallaxie/service-lens/internal/models"

	"github.com/katallaxie/pkg/dbx"
	"github.com/spf13/cobra"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gorm.Open(postgres.Open(config.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		store, err := dbx.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
		if err != nil {
			return err
		}

		return store.Migrate(
			cmd.Context(),
			&adapters.GothUser{},
			&adapters.GothAccount{},
			&adapters.GothSession{},
			&adapters.GothVerificationToken{},
			&models.Template{},
			&models.Ownable{},
			&models.Workflow{},
			&models.WorkflowState{},
			&models.WorkflowTransition{},
			&models.Workable{},
			&models.Reaction{},
			&models.ProfileQuestion{},
			&models.ProfileQuestionChoice{},
			&models.ProfileQuestionAnswer{},
			&models.Design{},
			&models.DesignRevision{},
			&models.DesignComment{},
			&models.DesignCommentRevision{},
			&models.Environment{},
			&models.Profile{},
			&models.Tag{},
			&models.Lens{},
			&models.Pillar{},
			&models.Question{},
			&models.Resource{},
			&models.Choice{},
			&models.Risk{},
			&models.Workload{},
			&models.WorkloadLensQuestionAnswer{},
			&models.Setting{},
		)
	},
}
