package service

import "slices"

import "fmt"

var ParamsArgs = []string{"-c", "create", "-h", "--help"}

func ShowHelp() {
	fmt.Println("Utilisation : chiS [paramètre] [optionnel]")
	fmt.Println("Paramètres disponibles :")
	for _, flag := range ParamsArgs {
		if flag == "-c" || flag == "create" {
			fmt.Println(flag,": " ,"create new project")
		}
		if flag == "-h" || flag == "--help" {
			fmt.Println(flag, ": ","show all command")
		}
	}
}

func IsAllowed(flag string) bool {
	return slices.Contains(ParamsArgs, flag)
}

