package main

import (
	"cursogo/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	fmt.Println(carro.Andar())
	fmt.Println("Resultado", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
