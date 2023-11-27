package error

import "fmt"

// Function to check an error
func IsError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
