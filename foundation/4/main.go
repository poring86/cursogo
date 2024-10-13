package main

import "fmt"

const a = "Hello"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Matheus"
	e float64 = 3.4
	f ID      = 1
)

func main() {

	fmt.Printf("O valor de E é %v \n", e)
	fmt.Printf("O tipo de E é %T \n", e)
	fmt.Printf("O tipo de F é %T \n", f)
	fmt.Println("O tipo de E é", e)
}
