package cmd

import (
	"fmt"
	"log"

	"github.com/clcosta/clcosta-cli/internal/storage"
	"github.com/clcosta/clcosta-cli/pkg/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "clcosta",
		Short: "clcosta: A utility CLI for various common tasks",
	}

	config.LoadEnviroment()

	pklStorage := storage.NewPklStorage()

	registerCommands(rootCmd, pklStorage)
	return rootCmd
}

func registerCommands(rootCmd *cobra.Command, pklStorage *storage.PklStorage) {
	pklConf, err := pklStorage.Load()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.LoadYamlConfig(pklConf.YAMLPathFile)
	if err != nil {
		fmt.Println("Error loading configuration file")
		log.Fatal(err)
	}

	rootCmd.AddCommand(newConfigureYAMLConfigCmd(config, pklStorage))
	rootCmd.AddCommand(newConfigureFilesCmd(config))
	rootCmd.AddCommand(newConfigureGitUserCmd(config))
	rootCmd.AddCommand(newConfigureSSHCmd(config))
}
