package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(fmt.Sprintf("%s must be a number only", prompt)) // Fixed error handling
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	if value2 == 0 {
		panic("Cannot divide by zero")
	}
	return value1 / value2
}

func getOperator() string {
	fmt.Print("Operator (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

// func main() {

// 	value1 := getInput("Enter value 1: ")
// 	value2 := getInput("Enter value 2: ")

// 	var result float64

// 	operator := getOperator()

// 	switch operator {
// 	case "+":
// 		result = add(value1, value2)
// 	case "-":
// 		result = subtract(value1, value2)
// 	case "*":
// 		result = multiply(value1, value2)
// 	case "/":
// 		result = divide(value1, value2)
// 	default:
// 		panic("Invalid operator")
// 	}

// 	fmt.Printf("Result: %.2f\n", result)
// }
