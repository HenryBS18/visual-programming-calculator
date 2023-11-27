package ui

import (
	"BasicCalculatorApp/calculator"
	"BasicCalculatorApp/error"
	"fmt"
)

func SquareRootMenu() {
	var a float64

	fmt.Print("Number: ")
	_, err := fmt.Scanln(&a)
	error.IsError(err)

	result, err2 := calculator.SquareRoot(a)
	error.IsError(err2)

	fmt.Println("Result is: ", result)
	fmt.Println("")
}
