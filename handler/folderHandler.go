package handler

import (
	"os"
	"path/filepath"
	"sync"
)

func CreateNecessaryFolder(directory string) {
	var wg sync.WaitGroup
	var folders = []string{filepath.Join(directory, "app"), filepath.Join(directory, "controller"), filepath.Join(directory, "repository"), filepath.Join(directory, "cmd"), filepath.Join(directory, "model"), filepath.Join(directory, "router")}
	for _, folder := range folders {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err := os.MkdirAll(f, 0755)
			if err != nil {
				ErrorHandler(err)
				return
			}
		}(folder)
	}

	wg.Wait()
}
