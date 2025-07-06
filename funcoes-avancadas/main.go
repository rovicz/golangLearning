package main

import "fmt"

// função de retornos nomeados.
func calcMath(a int, b int) (summ int, sub int) {
	summ = a + b
	sub = a - b

	return // apenas passe um retorno, já que o retorno já foi nomeado e definido.
}

func main() {
	a := 10
	b := 20

	// chamando a função e absorvendo seus retornos nomeados.
	summ, sub := calcMath(a, b)
	fmt.Println("Soma:", summ) // 30
	fmt.Println("Subtração:", sub) // -10
}