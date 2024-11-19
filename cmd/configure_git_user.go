package cmd

import (
	"github.com/clcosta/clcosta-cli/internal/git"
	"github.com/clcosta/clcosta-cli/pkg/config"

	"github.com/spf13/cobra"
)

func newConfigureGitUserCmd(config *config.ConfigYAML) *cobra.Command {
	var user string

	cmd := &cobra.Command{
		Use:   "gitUser [--user <type>]",
		Short: "Configure Git user locally",
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Configure(user)
		},
	}

	cmd.Flags().StringVarP(&user, "user", "u", "", "User type (e.g., personal, work)")
	return cmd
}
