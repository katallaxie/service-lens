package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

type ValidateCmdConfig struct {
	Schema   string
	Document string
}

var validateCfg = &ValidateCmdConfig{}

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration for errors",
	Long:  `Validate the configuration for errors`,
	RunE: func(cmd *cobra.Command, args []string) error {
		schema, err := os.ReadFile(validateCfg.Schema)
		if err != nil {
			return err
		}

		document, err := os.ReadFile(validateCfg.Document)
		if err != nil {
			return err
		}

		schemaLoader := gojsonschema.NewBytesLoader(schema)
		documentLoader := gojsonschema.NewBytesLoader(document)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			return err
		}

		if !result.Valid() {
			fmt.Printf("The document is not valid. see errors :\n")
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}

		return nil
	},
}
