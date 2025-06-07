// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	floatVal1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("The value should be number")
	}

	// Convert the second string to a float64
	floatVal2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("The value should be number")
	}

	// Calculate and return the result
	return floatVal1 + floatVal2
}

func main() {
	fmt.Println("--- Running Test Cases ---")

	// Example 1
	fmt.Println("\nExample 1:")
	val1_ex1 := "10.0"
	val2_ex1 := "5.5"
	result_ex1 := calculate(val1_ex1, val2_ex1)
	fmt.Printf("value1: %s\n", val1_ex1)
	fmt.Printf("value2: %s\n", val2_ex1)
	fmt.Printf("return: %f\n", result_ex1)

	// Example 2
	fmt.Println("\nExample 2:")
	val1_ex2 := "100.0"
	val2_ex2 := "-110.5"
	result_ex2 := calculate(val1_ex2, val2_ex2)
	fmt.Printf("value1: %s\n", val1_ex2)
	fmt.Printf("value2: %s\n", val2_ex2)
	fmt.Printf("return: %f\n", result_ex2)

	fmt.Println("\n--- Test Cases Finished ---")
}
