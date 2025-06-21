package main

import (
	"fmt"
	"os"

	"github.com/ptdrpg/chi/init/scrypt/helper"
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
	
	//request traitement
	if os.Args[1] == "-c" {
		projectName := os.Args[2]
		err := os.Mkdir(projectName, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
