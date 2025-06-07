// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	floatVal1 := convertInputToValue(input1)
	floatVal2 := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(floatVal1, floatVal2)
	case "-":
		return subtractValues(floatVal1, floatVal2)
	case "*":
		return multiplyValues(floatVal1, floatVal2)
	case "/":
		return divideValues(floatVal1, floatVal2)
	default:
		panic("Operation is not found")
	}
}

func convertInputToValue(input string) float64 {
	floatVal, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}

	return floatVal
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0.0 {
		panic("Can't divide by 0")
	}
	return value1 / value2
}

func main() {
	fmt.Println("--- Running Test Cases ---")

	fmt.Println("\nImage Example 1:")
	result1 := calculate("10.0", "5.5", "+")
	fmt.Printf("10.0 + 5.5 = %f\n", result1)

	fmt.Println("\nImage Example 2:")
	result2 := calculate("100.0", "2", "/")
	fmt.Printf("100.0 / 2 = %f\n", result2)

	// Test case from the text, designed to cause a panic
	fmt.Println("\nText Example (Testing Invalid Operation):")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Successfully caught expected panic: %v\n", r)
			fmt.Println("\n--- Test Cases Finished ---")
		}
	}()

	calculate("100", "2", "l")
}
