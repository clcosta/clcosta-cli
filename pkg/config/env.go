package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/clcosta/clcosta-cli/pkg/utils"
)

const (
	ConfigPathPKL    = "config.pkl"
	ConfigPathYAML   = "config.yaml"
	ConfigEnviroment = "CLCOSTA_PATH"
)

type Enviroment struct {
	BaseDir        string
	PklPath        string
	BaseYamlConfig string
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func gararanteePathExists(path string) {
	if !pathExists(path) {
		os.Mkdir(path, 0755)
	}
}

func LoadEnviroment() {
	currentOs := runtime.GOOS
	switch currentOs {
	case "windows":
		path := os.Getenv(ConfigEnviroment)
		if path == "" {
			os.Setenv(ConfigEnviroment, os.Getenv("USERPROFILE")+"/.clcosta")
		}
	case "linux":
		path := os.Getenv(ConfigEnviroment)
		if path == "" {
			os.Setenv(ConfigEnviroment, os.Getenv("HOME")+"/.clcosta")
		}
	}
}

func NewEnviromentConfig() *Enviroment {
	env := Enviroment{
		BaseDir:        os.Getenv(ConfigEnviroment),
		PklPath:        os.Getenv(ConfigEnviroment) + "/" + ConfigPathPKL,
		BaseYamlConfig: os.Getenv(ConfigEnviroment) + "/templates/" + ConfigPathYAML,
	}
	gararanteePathExists(env.BaseDir)
	templates := LoadTemplatesEmbed()

	if !pathExists(env.BaseYamlConfig) {
		gararanteePathExists(filepath.Dir(env.BaseYamlConfig))
		utils.WriteFile(env.BaseYamlConfig, templates["config"])
	}
	return &env
}
