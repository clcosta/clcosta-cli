package storage

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/clcosta/clcosta-cli/pkg/config"
)

type CurrentConf struct {
	YAMLPathFile string
}

type PklStorage struct {
	CurrentConf *CurrentConf
}

func NewPklStorage() *PklStorage {
	env := config.NewEnviromentConfig() // FIXME: get env by dependency injection
	if filepath.Base(env.BaseDir) == "bin" {
		env.BaseDir = filepath.Dir(env.BaseDir)
	}
	return &PklStorage{
		CurrentConf: &CurrentConf{YAMLPathFile: env.BaseYamlConfig},
	}
}

func (s *PklStorage) Save(pathFile string) error {
	env := config.NewEnviromentConfig() // FIXME: get env by dependency injection
	pklPath := env.PklPath

	fileContent, err := os.Create(pklPath)
	if err != nil {
		return err
	}
	defer fileContent.Close()

	encoder := gob.NewEncoder(fileContent)
	if err := encoder.Encode(CurrentConf{YAMLPathFile: pathFile}); err != nil {
		return err
	}
	return nil
}

func (s *PklStorage) Load() (*CurrentConf, error) {
	env := config.NewEnviromentConfig() // FIXME: get env by dependency injection
	pklPath := env.PklPath
	if _, err := os.Stat(pklPath); os.IsNotExist(err) {
		if err := s.Save(s.CurrentConf.YAMLPathFile); err != nil {
			return nil, err
		}
	}

	var currentConf CurrentConf
	fileContent, err := os.Open(pklPath)
	if err != nil {
		err := fmt.Errorf("error opening config.pkl file: %v", err)
		return nil, err
	}
	defer fileContent.Close()

	fileInfo, err := fileContent.Stat()
	if err != nil {
		err := fmt.Errorf("error getting file info for config.pkl: %v", err)
		return nil, err
	}
	if fileInfo.Size() == 0 {
		// if the file it's empty then write the default values
		if err := s.Save(s.CurrentConf.YAMLPathFile); err != nil {
			err := fmt.Errorf("error saving default values to config.pkl file: %v", err)
			return nil, err
		}
	}

	decoder := gob.NewDecoder(fileContent)
	if err := decoder.Decode(&currentConf); err != nil {
		err := fmt.Errorf("error decoding config.pkl file: %v", err)
		return nil, err
	}
	return &currentConf, nil
}
