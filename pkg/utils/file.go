package utils

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}

	fmt.Printf("File created successfully: %s\n", filePath)
	return nil
}

func ReadFileContent(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info for file %s: %w", filePath, err)
	}

	fileSize := fileInfo.Size()
	fileContent := make([]byte, fileSize)

	_, err = file.Read(fileContent)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return fileContent, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
