package main

import "fmt"

func main() {
	fmt.Println(sum(2, 4, 5, 12, 10, 7))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
