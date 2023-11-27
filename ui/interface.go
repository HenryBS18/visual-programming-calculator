package ui

import (
	"BasicCalculatorApp/error"
	ui "BasicCalculatorApp/ui/menu"
	"fmt"
)

// Main Program
func Calculator() {
	for {
		var menu string

		fmt.Println("==================")
		fmt.Println("--- Calculator ---")
		fmt.Println("==================")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Power Of")
		fmt.Println("6. Square Root")
		fmt.Println("7. Exit")

		fmt.Print("Menu: ")
		_, err := fmt.Scanln(&menu)
		error.IsError(err)

		switch menu {
		case "1":
			ui.AddMenu()
		case "2":
			ui.SubtractMenu()
		case "3":
			ui.MultiplyMenu()
		case "4":
			ui.DivideMenu()
		case "5":
			ui.PowerOfMenu()
		case "6":
			ui.SquareRootMenu()
		case "7":
			return
		}
	}
}
