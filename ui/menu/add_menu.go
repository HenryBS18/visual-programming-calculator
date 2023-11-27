package ui

import (
	"BasicCalculatorApp/calculator"
	"BasicCalculatorApp/error"
	"fmt"
)

func AddMenu() {
	var a float64
	var b float64

	fmt.Print("Number: ")
	_, err := fmt.Scanln(&a)
	error.IsError(err)

	fmt.Print("Add by: ")
	_, err2 := fmt.Scanln(&b)
	error.IsError(err2)

	result := calculator.Add(a, b)

	fmt.Println("Result is: ", result)
	fmt.Println("")
}
