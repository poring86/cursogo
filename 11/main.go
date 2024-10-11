package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

type Pessoa interface {
	Desativar()
}

func main() {
	matheus := Client{
		Nome:  "Matheus",
		Idade: 30,
		Ativo: true,
	}
	matheus.Ativo = false

	matheus.Estado = "MS"
	matheus.Endereco.Numero = 20

	matheus.Desativar()

	fmt.Println(matheus.Ativo)

	fmt.Println(matheus.Endereco)
	fmt.Println(matheus.Endereco.Estado)

	fmt.Println(matheus.Numero)
}
