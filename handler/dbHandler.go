package handler

import "github.com/AlecAivazis/survey/v2"

func AskDB() string {
	var db string
	prompt := &survey.Select{
		Message: "Choisis ta base de données:",
		Options: []string{"postgres", "mysql", "sqlite"},
	}
	survey.AskOne(prompt, &db)
	return db
}
