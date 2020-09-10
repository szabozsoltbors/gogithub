package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

// ListDir - is a function wich returns an array
// with all the filenames from a directory
func ListDir(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	return files
}
