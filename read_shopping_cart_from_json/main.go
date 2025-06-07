// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	carts := make([]cartItem, 0)

	// Your code goes here.
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var item cartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		carts = append(carts, item)
	}

	return carts
}

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
  		 {"name":"orange","price":1.50,"quantity":8},
  		 {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	fmt.Println("--- Running Test Case ---")
	fmt.Println("Successfully parsed JSON into the following slice of structs:")
	fmt.Println(result)
	fmt.Println("\n--- Test Case Finished ---")
}
