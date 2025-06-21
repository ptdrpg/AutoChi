package handler

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func AddAllDependancies(directory string, db string) {
	var wg sync.WaitGroup
	var ressources = []string{"gorm.io/gorm", "github.com/go-chi/chi/v5", "github.com/joho/godotenv", filepath.Join("gorm.io/driver/", db)}
	for _, dep := range ressources {
		wg.Add(1)
		go func(d string, dir string) {
			defer wg.Done()
			installDep := exec.Command("go", "get", d)
			installDep.Dir = dir
			installDep.Stderr = os.Stderr
			err := installDep.Run()
			if err != nil {
				ErrorHandler(err)
				return
			}
		}(dep, directory)
	}

	wg.Wait()
}
