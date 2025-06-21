package templates

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteMain(directory string) error {
	content := fmt.Sprintf(`
		package main

import "%s/cmd"

func main() {
	cmd.Execute()
}

	`, directory)

	err := os.WriteFile(filepath.Join(directory, "main.go"), []byte(content),0755)
	if err != nil {
		return err
	}
	return nil
}
