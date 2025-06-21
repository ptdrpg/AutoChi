package service

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ptdrpg/chi/init/scrypt/handler"
)

//create project's base
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

	//create all necessary folder
	handler.CreateNecessaryFolder(directory)

	//Get all dependancies
	handler.AddAllDependancies(directory, db)

	//create main.go
	mainPath:= filepath.Join(directory, "main.go")
	file, err := os.Create(mainPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	file.Close()

}
