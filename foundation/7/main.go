package main

import "fmt"

func main() {
	salarios := map[string]int{"Matheus": 1000, "Junior": 2000, "Laura": 3000}
	fmt.Println(salarios["Matheus"])
	delete(salarios, "Matheus")
	salarios["Matt"] = 5000
	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
