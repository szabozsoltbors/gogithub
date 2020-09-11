package util

import (
	"os"
	"path/filepath"
)

// ListDir - returns an array with all the filenames from a directory
func ListDir(root string) ([]string, error) {

	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
