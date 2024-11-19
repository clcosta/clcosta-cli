package config

import (
	"errors"

	"github.com/clcosta/clcosta-cli/pkg/utils"
	"gopkg.in/yaml.v3"
)

type GitAccount struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

type SSHConf struct {
	Key string `yaml:"key"`
}

type FileConf struct {
	Active   bool   `yaml:"active"`
	Template string `yaml:"template"`
}

type ConfigYAML struct {
	Git       map[string][]GitAccount `yaml:"git"`
	SSH       map[string][]SSHConf    `yaml:"ssh"`
	ConfFiles map[string][]FileConf   `yaml:"confFiles"`
}

var (
	ErrInvalidConfigGit       = func(account string) error { return errors.New("invalid configuration for git account: " + account) }
	ErrInvalidConfigSSH       = func(key string) error { return errors.New("invalid SSH configuration for key: " + key) }
	ErrInvalidConfigConfFiles = func(key string) error { return errors.New("invalid configuration for file key: " + key) }
	ErrInvalidConfig          = func(key string) error { return errors.New("invalid configuration for key: " + key) }
)

func LoadYamlConfig(filePath string) (*ConfigYAML, error) {
	fileContent, err := utils.ReadFileContent(filePath)
	if err != nil {
		return nil, err
	}

	var config ConfigYAML
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *ConfigYAML) Validate() error {
	// for each account in git check if the email and user are not empty
	if validateGitAccounts(c.Git) != nil {
		return ErrInvalidConfig("git")
	}
	// for each shh key configured check if the key is not empty
	if validateSSHKeys(c.SSH) != nil {
		return ErrInvalidConfig("ssh")
	}
	// for each file configuration check if the template exists and the content is not empty
	if validateConfFiles(c.ConfFiles) != nil {
		return ErrInvalidConfig("confFiles")
	}
	return nil
}

func validateGitAccounts(accounts map[string][]GitAccount) error {
	for account, gitAccounts := range accounts {
		if len(gitAccounts) == 0 {
			return ErrInvalidConfigGit(account)
		}
	}
	return nil
}

func validateSSHKeys(sshKeys map[string][]SSHConf) error {
	for ssh, sshConf := range sshKeys {
		if len(sshConf) == 0 {
			return ErrInvalidConfigSSH(ssh)
		}
	}
	return nil
}

func validateConfFiles(confFiles map[string][]FileConf) error {
	for file, fileConf := range confFiles {
		if len(fileConf) == 0 {
			return ErrInvalidConfigConfFiles(file)
		}
		println("File:" + file)
	}
	return nil
}
