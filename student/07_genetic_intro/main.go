package main

import "fmt"

func fitness(candidate string) int {
	count := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == '1' {
			count++
		}
	}
	return count
}

func main() {
	a, b := "1001", "0110"

	// Crossover after character 2 (index 2): take first 2 from a, rest from b
	child := a[:2] + b[2:]
	fmt.Printf("Child after crossover: %s  fitness = %d\n", child, fitness(child))

	// Mutate second character (index 1) to '1'
	childBytes := []byte(child)
	childBytes[1] = '1'
	child = string(childBytes)
	fmt.Printf("Child after mutation:  %s  fitness = %d\n", child, fitness(child))
}
