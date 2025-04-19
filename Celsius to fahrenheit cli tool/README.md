# Project Documentation: Temperature Conversion CLI Tool
## Overview
The Temperature Conversion CLI Tool is a beginner-friendly command-line application designed to convert Celsius to Fahrenheit. This project demonstrates foundational programming concepts in Go, such as input handling, calculations, and formatted output.

## Key Features
1. **User Input**: Accepts a temperature in Celsius from the user.
2. **Conversion Logic**: Uses a simple mathematical formula to convert Celsius to Fahrenheit.
3. **Formatted Output**: Displays the result with precision.

## Motivation 
The primary goal of this project was to learn the basics of the Go programming language while documenting the step-by-step process to arrive at the solution. This approach helps solidify understanding and provides a resource for others who may be new to Go or coding in general.

## Code Example
```Go
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
```

## How it Works
1. The program prompts the user to enter a temperature in Celsius.
2. It reads the input using `fmt.Scanln` and stores it in a variable.
3. The conversion formula `(Celsius * 9/5) + 32` is applied to calculate Fahrenheit.
4. The result is deisplayed using `fmt.Printf` for neat formatting.

## Why This Project is Important
1. **Learning Go Basics**: Builds a solid foundation in Go by practicing essential programming concepts like input/output, arithmetic operations, and data types.
2. **Problem Solving**: Converts a real-world scenario (temperature conversion) into a simple software solution, showcasing practical problem-solving skills.
3. **Portfolio Value**: Demonstrates clean, functional code in Go, a language valued for its efficiency and scalability.
4. **Expandability**: Serves as a stepping stone for more complex projects, such as interactive CLI tools or backend services.

## References for Further Learning
To help beginners understand the concepts used in this project, here are some resources:
- [The Go Programming Language Tutorial](https://go.dev/doc/tutorial/): A comprehensive introduction to the Go programming language, perfect for beginners.
- [Go by Example](https://gobyexample.com/): A collection of concise examples for learning various Go programming concepts, from basics to advanced.
- [Temperature Conversion Formula](https://www.mathsisfun.com/temperature-conversion.html): To help beginners understand the concepts used in this project, here are some resources:
- [Official Go Documentation](https://pkg.go.dev/fmt): Detailed documentation for Go's `fmt` package, used for input and output in this project.
