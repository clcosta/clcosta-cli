package config

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/// This file will embed the templates files for configuration works

var embeddedFiles embed.FS

var (
	paths = []string{
		"config.yaml",
		".editorconfig",
		".pre-commit-config.yaml",
	}
)

func LoadTemplatesEmbed() map[string][]byte {
	templateMap := make(map[string][]byte)
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return templateMap
	}
	execDir := filepath.Dir(execPath)

	for _, path := range paths {
		absPath := filepath.Join(execDir, "templates", path)
		data, err := os.ReadFile(absPath)
		if err != nil {
			fmt.Println("Error reading file: ", absPath, err)
			continue
		}
		templateMap[normalizePathKey(path)] = data
	}

	return templateMap
}

func normalizePathKey(path string) string {
	path = strings.Replace(path, ".yaml", "", -1)
	re := regexp.MustCompile(`[0-9\p{P}]`)
	path = re.ReplaceAllString(path, "")
	return path
}
