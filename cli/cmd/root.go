package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	system  = "unknown"
	arch    = "unknown"
)

// Config ...
type Config struct {
	Flags *Flags
}

type Flags struct {
}

var versionFmt = fmt.Sprintf("%s-%s (%s) %s/%s", version, commit, date, system, arch)

func Init() error {
	ctx := context.Background()

	ValidateCmd.Flags().StringVarP(&validateCfg.Schema, "schema", "s", "", "Path to the schema file")
	ValidateCmd.Flags().StringVarP(&validateCfg.Document, "document", "d", "", "Path to the document file")

	Root.AddCommand(ValidateCmd)

	Root.SilenceUsage = true

	err := Root.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

var Root = &cobra.Command{
	Use:     "cli",
	Short:   "A tool to help interface with the service lens",
	Long:    `A tool to help interface with the service lens`,
	Version: versionFmt,
}
