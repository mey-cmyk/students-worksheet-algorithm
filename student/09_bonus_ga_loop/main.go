package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func fitness(candidate string) int {
	count := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == '1' {
			count++
		}
	}
	return count
}

func randomCandidate() string {
	var b strings.Builder
	for i := 0; i < 4; i++ {
		if rand.Intn(2) == 1 {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
	}
	return b.String()
}

func crossover(a, b string) string {
	cut := rand.Intn(3) + 1 // cut after 1, 2, or 3 bits
	return a[:cut] + b[cut:]
}

func mutate(candidate string) string {
	b := []byte(candidate)
	idx := rand.Intn(4)
	if b[idx] == '0' {
		b[idx] = '1'
	} else {
		b[idx] = '0'
	}
	return string(b)
}

func tournamentSelect(pop []string) string {
	// Pick 2 random candidates, return the better one
	i, j := rand.Intn(len(pop)), rand.Intn(len(pop))
	if fitness(pop[i]) >= fitness(pop[j]) {
		return pop[i]
	}
	return pop[j]
}

func bestFitness(pop []string) int {
	best := 0
	for _, c := range pop {
		f := fitness(c)
		if f > best {
			best = f
		}
	}
	return best
}

func main() {
	rand.Seed(time.Now().UnixNano())

	popSize := 10
	maxGenerations := 50
	mutationRate := 0.3

	// Initialize population
	pop := make([]string, popSize)
	for i := 0; i < popSize; i++ {
		pop[i] = randomCandidate()
	}

	for gen := 1; gen <= maxGenerations; gen++ {
		// Print best fitness this generation
		best := bestFitness(pop)
		fmt.Printf("Generation %d: best fitness = %d\n", gen, best)

		if best == 4 {
			fmt.Println("Found optimal solution (1111)!")
			break
		}

		// Create next generation
		newPop := make([]string, popSize)
		for i := 0; i < popSize; i += 2 {
			parent1 := tournamentSelect(pop)
			parent2 := tournamentSelect(pop)

			child1 := crossover(parent1, parent2)
			child2 := crossover(parent2, parent1)

			if rand.Float64() < mutationRate {
				child1 = mutate(child1)
			}
			if rand.Float64() < mutationRate {
				child2 = mutate(child2)
			}

			newPop[i] = child1
			if i+1 < popSize {
				newPop[i+1] = child2
			}
		}
		pop = newPop
	}
}
