package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados do arquivo!"))
	// tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)

		println(n)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
