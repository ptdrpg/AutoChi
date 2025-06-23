package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ptdrpg/chiS/handler"
)

func WriteMain(directory string) {
	content := fmt.Sprintf(`
		package main

import "%s/cmd"

func main() {
	cmd.Execute()
}

	`, directory)

	err := os.WriteFile(filepath.Join(directory, "main.go"), []byte(content),0755)
	if err != nil {
		handler.ErrorHandler(err)
	}
}
