package main

import "fmt"

func main() {
	// maneira pratica de criação de variáveis e tipagem.
	var nome string = "Victor"
	var idade int = 20

	// maneira mais inteligente de criação de variáveis e tipagem.
	cidade := "Goiânia"

	// outra maneira de criação de variáveis em maior escala.
	var (
		cat string = "Chloe"
		catAge int = 1
	)

	fmt.Println(nome, "is", idade, "years old and lives in", cidade, "he got a cat named", cat, "and she is", catAge, "years old")
}