package main

import "fmt"

// Write your answer here, and then test your code.

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	price := 100.0 / float64(len(items))

	// Your code goes here
	for _, v := range items {
		result[v] = price
	}
	return result
}

func main() {
	fmt.Println("--- Running Test Cases ---")

	// Example 1
	fmt.Println("\nExample 1:")
	input1 := []string{"apple", "orange"}
	fmt.Printf("Input: %v\n", input1)
	fmt.Println("Return:")
	result1 := convertToMap(input1)
	// Note: Map iteration order is not guaranteed, but all key-value pairs will be present.
	for key, value := range result1 {
		fmt.Printf("%s: %f\n", key, value)
	}

	// Example 2
	fmt.Println("\nExample 2:")
	input2 := []string{"banana", "orange", "apple"}
	fmt.Printf("Input: %v\n", input2)
	fmt.Println("Return:")
	result2 := convertToMap(input2)
	for key, value := range result2 {
		fmt.Printf("%s: %f\n", key, value)
	}

	fmt.Println("\n--- Test Cases Finished ---")
}
