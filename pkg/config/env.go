package config

import (
	"os"
	"runtime"
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

func LoadEnviroment() {
	currentOs := runtime.GOOS
	switch currentOs {
	case "windows":
		path := os.Getenv(ConfigEnviroment)
		if path == "" {
			os.Setenv(ConfigEnviroment, os.Getenv("USERPROFILE")+"/.clcosta")
			if _, err := os.Stat(os.Getenv(ConfigEnviroment)); os.IsNotExist(err) {
				os.Mkdir(os.Getenv(ConfigEnviroment), 0755)
			}
		}
	case "linux":
		path := os.Getenv(ConfigEnviroment)
		if path == "" {
			os.Setenv(ConfigEnviroment, os.Getenv("HOME")+"/.clcosta")
			if _, err := os.Stat(os.Getenv(ConfigEnviroment)); os.IsNotExist(err) {
				os.Mkdir(os.Getenv(ConfigEnviroment), 0755)
			}
		}
	}
}

func NewEnviromentConfig() *Enviroment {
	return &Enviroment{
		BaseDir:        os.Getenv(ConfigEnviroment),
		PklPath:        os.Getenv(ConfigEnviroment) + "/" + ConfigPathPKL,
		BaseYamlConfig: os.Getenv(ConfigEnviroment) + "/templates/" + ConfigPathYAML,
	}
}
