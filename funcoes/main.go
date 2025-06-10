package main

import (
	"fmt"
)

func Somar(a int, b int) int {
	return a + b
}

func main() {
	var a int = 10
	var b int = 20
	var soma int
	
	var dividir = func(a int) int {
		return a / 2
	}

	// pode-se fazer uma função de mais de um retorno (?)
	var multiplicarELevarAPotencia = func(a int) (int, int) {
		return a * 2, a * a // pelo amor de deus, o que é isso?
	} // i love it.

	fmt.Println("O resultado da soma é:", Somar(a, b))
	soma = Somar(a, b)

	fmt.Println("O resultado da divisão é:", dividir(soma))

	// extrai os dois resultados em duas variáveis diferentes.
	resultadoMultiplicacao, resultadoPotencia := multiplicarELevarAPotencia(soma) // caso não queira usar um dos valores do retorno, pode-se usar o underscore (_).
	fmt.Println("O Resultado da multiplicação por 2 é", resultadoMultiplicacao, "e o resultado da potência é:", resultadoPotencia)
}