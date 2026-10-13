package main

import "fmt"

func main() {
	start, births, arrivals, deaths, departures := 1000, 20, 5, 10, 4
	end := start + births + arrivals - deaths - departures
	fmt.Printf("Population: %d\n", end)
}
