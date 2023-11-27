package ui

import (
	"fmt"
	"time"
)

// Main function
func Goroutines() {
	go p()
	go q()
	r()
	time.Sleep(1 * time.Second)
}

func p() {
	result := 1000000 * 1000000 * 1000000
	fmt.Println(result)
}

func q() {
	result := 1 + 1
	fmt.Println(result)
}

func r() {
	result := 5 - 2
	fmt.Println(result)
}
