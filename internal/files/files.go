package files

import (
	"github.com/clcosta/clcosta-cli/pkg/config"
	"github.com/clcosta/clcosta-cli/pkg/utils"
)

func CreateFiles(config *config.ConfigYAML) error {
	for _, files := range config.ConfFiles {
		for _, file := range files {
			if file.Active {
				// create the file with the template content
				templateContent, err := utils.ReadFileContent(file.Template)
				if err != nil {
					return err
				}
				err = utils.WriteFile(file.Template, templateContent)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
