package main

import (
	"fmt"
)

func main() {
	// Prompt user for their input
	fmt.Println("Enter temperature in Celsius:")

	// Declare variable to store user input
	var Celsius float64
	fmt.Scanln(&Celsius)

	// Convert Celsius to Fahrenheit
	Fahrenheit := (Celsius * 9 / 5) + 32

	// Display the result
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", Fahrenheit)
}
