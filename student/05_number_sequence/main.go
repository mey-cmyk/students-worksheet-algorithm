package main

import "fmt"

func main() {
	first, second := 5, 3

	// Without a loop
	third := first + second
	fourth := second + third
	fifth := third + fourth

	fmt.Printf("%d %d %d\n", third, fourth, fifth)

	// Optional: with a for loop
	fmt.Println("--- With loop ---")
	a, b := 1, 1 // Example: starting values 1 1
	for i := 0; i < 3; i++ {
		c := a + b
		fmt.Printf("%d ", c)
		a, b = b, c
	}
	fmt.Println()
}
