package cmd

import (
	"github.com/clcosta/clcosta-cli/internal/storage"
	"github.com/clcosta/clcosta-cli/internal/yaml_config"
	"github.com/clcosta/clcosta-cli/pkg/config"
	"github.com/spf13/cobra"
)

func newConfigureYAMLConfigCmd(config *config.ConfigYAML, pklStorage *storage.PklStorage) *cobra.Command {
	var configPath string

	cmd := &cobra.Command{
		Use:   "config [--config <path>]",
		Short: "Set the YAML configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return yaml_config.ConfigureYAML(config, pklStorage, configPath)
		},
	}

	cmd.Flags().StringVarP(&configPath, "config", "c", "", "Config Path to YAML file")
	return cmd
}
