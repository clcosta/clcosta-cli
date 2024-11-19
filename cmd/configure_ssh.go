package cmd

import (
	"github.com/clcosta/clcosta-cli/internal/ssh"
	"github.com/clcosta/clcosta-cli/pkg/config"

	"github.com/spf13/cobra"
)

func newConfigureSSHCmd(config *config.ConfigYAML) *cobra.Command {
	return &cobra.Command{
		Use:   "configureSSH",
		Short: "Set up SSH configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ssh.Configure()
		},
	}
}
