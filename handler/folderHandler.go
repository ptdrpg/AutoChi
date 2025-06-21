package handler

import (
	"os"
	"path/filepath"
)

func CreateNecessaryFolder(directory string) {
	var folders = []string{filepath.Join(directory, "app"), filepath.Join(directory, "controller"), filepath.Join(directory, "repository"), filepath.Join(directory, "cmd"), filepath.Join(directory, "model"), filepath.Join(directory, "router")}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			ErrorHandler(err)
			return
		}
	}
}
