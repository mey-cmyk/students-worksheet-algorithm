package main

import "fmt"

func main() {
	unitPrice := 18000
	quantity := 2
	discount := 5000

	subtotal := unitPrice * quantity
	total := subtotal - discount

	fmt.Printf("Subtotal: %d\n", subtotal)
	fmt.Printf("Total: %d\n", total)

	// Change quantity to 3 and recompute
	quantity = 3
	subtotal = unitPrice * quantity
	total = subtotal - discount

	fmt.Printf("Subtotal: %d\n", subtotal)
	fmt.Printf("Total: %d\n", total)
}
