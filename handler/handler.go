package handler

import (
	"log"
)

//handle erro
func ErrorHandler(props error) {
	log.Fatal(props.Error())
}