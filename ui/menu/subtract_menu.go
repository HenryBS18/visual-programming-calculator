package ui

import (
	"BasicCalculatorApp/calculator"
	"BasicCalculatorApp/error"
	"fmt"
)

func SubtractMenu() {
	var a float64
	var b float64

	fmt.Print("Number: ")
	_, err := fmt.Scanln(&a)
	error.IsError(err)

	fmt.Print("Subtract by: ")
	_, err2 := fmt.Scanln(&b)
	error.IsError(err2)

	result := calculator.Subtract(a, b)

	fmt.Println("Result is: ", result)
	fmt.Println("")
}
