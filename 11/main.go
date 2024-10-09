package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	matheus := Client{
		Nome:  "Matheus",
		Idade: 30,
		Ativo: true,
	}
	matheus.Ativo = false

	fmt.Println(matheus.Ativo)
}
