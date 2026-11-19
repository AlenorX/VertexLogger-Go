package main

import (
	"fmt"
	"vertexlogger/logger"
	_ "vertexlogger/testing"
)

// Here checking job my logger on Go

var source string = "Kekus"

func main() {
	/*
		text, err := testing.CheckLogger(source, logger.Info)
		if err != nil {
			fmt.Println("Возникла ошибка во время тестирования")
		}
		fmt.Println(text)
	*/

	output, err := logger.Info(source)
	if err != nil {
		fmt.Println("Возникла ошибка в logger")
	}
	fmt.Println(output)
}
