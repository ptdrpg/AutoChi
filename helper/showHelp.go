package helper

import "slices"

import "fmt"

var ParamsArgs = []string{"-c", "create", "-h", "--help"}

func ShowHelp() {
	fmt.Println("Utilisation : ./test [paramètre] [optionnel]")
	fmt.Println("Paramètres disponibles :")
	for _, flag := range ParamsArgs {
		fmt.Println("  ", flag)
	}
}

func IsAllowed(flag string) bool {
	return slices.Contains(ParamsArgs, flag)
}

