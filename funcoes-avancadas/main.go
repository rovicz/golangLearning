package main

import "fmt"

// função de retornos nomeados.
func calcMath(a int, b int) (summ int, sub int) {
	summ = a + b
	sub = a - b

	return // apenas passe um retorno, já que o retorno já foi nomeado e definido.
}

// função variática.
func onlySumm(numbers ...int) int { // passando um type seguindo o formato de "...type" declara esse argumento como variável, sendo limitado a UMA variável, sempre como o último argumento da função.
	total := 0

	// realizando uma soma dentro do loop.
	for _, number := range numbers {
		total += number
	}

	fmt.Println("Os argumentos são:", numbers)
	return total
}

func main() {
	a := 10
	b := 20

	// chamando a função e absorvendo seus retornos nomeados.
	summ, sub := calcMath(a, b)
	fmt.Println("Soma:", summ) // 30
	fmt.Println("Subtração:", sub) // -10

	// chamando a função variática e passando 10 argumentos.
	summVariaty := onlySumm(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // 10 argumentos.
	fmt.Println("Soma da função variática:", summVariaty) // 55

	// função anônima.
	func () {
		fmt.Println("Olá mundo!")
	}()

	// função anônima (com parametros).
	func (text string, numbers ...int) {
		total := 0

		for _, number := range numbers {
			total += number
		}

		fmt.Println(text, total)
	}("Total da soma via função anônima:", 1, 10, 15, 50, 100)
}