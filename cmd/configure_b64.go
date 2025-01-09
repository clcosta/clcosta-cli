package cmd

import (
	"github.com/clcosta/clcosta-cli/internal/b64"
	"github.com/spf13/cobra"
)

func newConfigureB64Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "b64 [string] [ -d decode]",
		Short: "Base64 encode/decode strings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			decode, _ := cmd.Flags().GetBool("decode")
			return b64.Configure(args[0], decode)
		},
	}

	cmd.Flags().BoolP("decode", "d", false, "Decode the string")

	return cmd
}
