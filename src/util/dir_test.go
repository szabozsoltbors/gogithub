package util

import (
	"testing"
)

func TestListDir(t *testing.T) {

	files, err := ListDir("./")

	if err != nil {
		t.Error("Unable to load the HTML templates")
	}

	result := false
	for _, file := range files {
		if file == "dir_test.go" {
			result = true
		}
	}

	if !result {
		t.Error("Unable to find dir_test.go")
	}

}
