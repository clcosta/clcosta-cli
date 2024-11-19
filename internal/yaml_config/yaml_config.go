package yaml_config

import (
	"github.com/clcosta/clcosta-cli/internal/storage"
	configPkg "github.com/clcosta/clcosta-cli/pkg/config"
	"github.com/clcosta/clcosta-cli/pkg/utils"
)

func ConfigureYAML(config *configPkg.ConfigYAML, pklStorage *storage.PklStorage, newConfigPath string) error {
	env := configPkg.NewEnviromentConfig()
	if newConfigPath == "" {
		templateContent, err := utils.ReadFileContent(env.BaseYamlConfig)
		if err != nil {
			return err
		}
		utils.WriteFile(newConfigPath, templateContent)
	}

	newConfig, err := configPkg.LoadYamlConfig(newConfigPath)
	if err != nil {
		return err
	}
	err = newConfig.Validate()
	if err != nil {
		return err
	}

	pklStorage.Save(newConfigPath)
	return nil
}
