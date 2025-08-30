package internal

import (
	"os"
	"path/filepath"
)

func GetFileList(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

var ownerRead = os.FileMode(0400)
var ownerWrite = os.FileMode(0200)
var groupRead = os.FileMode(0040)
var othersRead = os.FileMode(0004)

func WriteFile(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), ownerRead|ownerWrite|groupRead|othersRead)
}

func ChangeExtension(filePath string, newExt string) string {
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)
	name := base[:len(base)-len(filepath.Ext(base))]
	newPath := filepath.Join(dir, name+newExt)
	return newPath
}
