package service

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ptdrpg/chiS/handler"
	"github.com/ptdrpg/chiS/templates"
)

// create project's base
func Create(directory string, db string) {
	//create project folder
	err := os.Mkdir(directory, 0755)
	if err != nil {
		handler.ErrorHandler(err)
	}

	//init go inside project
	initGo := exec.Command("go", "mod", "init", directory)
	initGo.Dir = directory
	initGo.Stderr = os.Stderr
	err = initGo.Run()
	if err != nil {
		handler.ErrorHandler(err)
	}

	initGit := exec.Command("git", "init")
	initGit.Dir = directory
	initGit.Stderr = os.Stderr
	err = initGit.Run()
	if err != nil {
		handler.ErrorHandler(err)
	}

	//create all necessary folder
	handler.CreateNecessaryFolder(directory)
	//Get all dependancies
	handler.AddAllDependancies(directory, db)

	//create main.go
	mainPath := filepath.Join(directory, "main.go")
	main, err := os.Create(mainPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	main.Close()

	templates.ConfigDB(directory, db)
	templates.WriteMain(directory)
	templates.WritteCmd(directory)
	templates.WriteRepository(directory)
	templates.WriteController(directory)
	templates.WritteRouter(directory)
}
