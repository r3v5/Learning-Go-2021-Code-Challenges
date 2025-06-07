// Write your answer here, and then test your code.

package main

import "fmt"

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	result := 0.0

	for i := 0; i < len(cart); i++ {
		result = result + (cart[i].price * float64(cart[i].quantity))
	}

	return result
}

func main() {

	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}

	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)

	fmt.Println("--- Running Test Case ---")
	fmt.Printf("Calculating total for %d item types in the cart...\n", len(cart))
	fmt.Printf("Total cart value: %f\n", result)
	fmt.Println("--- Test Case Finished ---")
}
