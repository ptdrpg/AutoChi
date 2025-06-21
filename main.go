package main

import (
	"fmt"
	"os"

	"github.com/ptdrpg/chi/init/scrypt/handler"
	"github.com/ptdrpg/chi/init/scrypt/helper"
	"github.com/ptdrpg/chi/init/scrypt/service"
)

func main() {

	if len(os.Args) < 3 {
		helper.ShowHelp()
	}

	//flag check if allowed
	if !helper.IsAllowed(os.Args[1]) {
		fmt.Printf("flag %s not found", os.Args[1])
		helper.ShowHelp()
		return
	}

	if os.Args[1] == "-c" {
		projectName := os.Args[2]
		db := handler.AskDB()
		service.Create(projectName, db)
	}

}
