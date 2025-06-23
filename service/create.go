package service

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ptdrpg/chi/init/scrypt/handler"
	"github.com/ptdrpg/chi/init/scrypt/templates"
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

	//create all necessary folder
	handler.CreateNecessaryFolder(directory)
	//Get all dependancies
	handler.AddAllDependancies(directory, db)

	variableEnv := `DB_HOST=localhost
	DB_USER=dbusername
	DB_PASSWORD=dbpassword
	DB_NAME=dbname
	DB_PORT=5432
	SECRET_KEY=secretkey`
	envPath := filepath.Join(directory, ".env")
	env, err := os.Create(envPath)
	if err != nil {
		handler.ErrorHandler(err)
	}
	env.Write([]byte(variableEnv))
	env.Close()

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
