package main

import "fmt"

func main() {
	var minhaVar interface{} = "Matheus"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v \n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v \n", res2)
}
