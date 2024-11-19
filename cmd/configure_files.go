package cmd

import (
	"github.com/clcosta/clcosta-cli/internal/files"
	"github.com/clcosta/clcosta-cli/pkg/config"
	"github.com/spf13/cobra"
)

func newConfigureFilesCmd(config *config.ConfigYAML) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "files",
		Short: "Create the configurated files in current directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return files.CreateFiles(config)
		},
	}

	return cmd
}
