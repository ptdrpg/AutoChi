package handler

import (
	"fmt"
)

//handle error
func ErrorHandler(props error) {
	fmt.Println(props.Error())
}