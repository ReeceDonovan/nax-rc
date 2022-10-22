package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("error checking if path %s exists: %w", path, err)
}

func CheckPathIsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("error checking if path %s is a directory: %w", path, err)
	}
	return info.IsDir(), nil
}

func CreateDirectory(dir string) error {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", dir, err)
	}
	return nil
}

func CreateDirectories(dirs []string) error {
	for _, dir := range dirs {
		err := CreateDirectory(dir)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}
	return nil
}

func GetFilesInDirectory(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}

func CreateFile(file string) error {
	_, err := os.Create(file)
	return err
}

func CreateFiles(files []string) error {
	for _, file := range files {
		err := CreateFile(file)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", file, err)
		}
	}
	return nil
}

func WriteToFile(file string, data string) error {
	return os.WriteFile(file, []byte(data), os.ModePerm)
}

func ClearFiles(paths []string) error {
	for _, path := range paths {
		err := CreateFile(path)
		if err != nil {
			return fmt.Errorf("error clearing file %s: %w", path, err)
		}
	}
	return nil
}

func GetFileModificationTime(file string, formatString string) (string, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", fmt.Errorf("error getting file modification time: %w", err)
	}
	return info.ModTime().Format(formatString), nil
}

func GetFileContent(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("error getting file content: %w", err)
	}
	return string(content), nil
}
