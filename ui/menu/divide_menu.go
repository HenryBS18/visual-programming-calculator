package ui

import (
	"BasicCalculatorApp/calculator"
	"BasicCalculatorApp/error"
	"fmt"
)

func DivideMenu() {
	var a float64
	var b float64

	fmt.Print("Number: ")
	_, err := fmt.Scanln(&a)
	error.IsError(err)

	fmt.Print("Divided by: ")
	_, err2 := fmt.Scanln(&b)
	error.IsError(err2)

	result, err3 := calculator.Divide(a, b)
	error.IsError(err3)

	fmt.Println("Result is: ", result)
	fmt.Println("")
}
