package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		fmt.Println("invalid input")
		return
	}

	fmt.Println(y%x == 0)
}
