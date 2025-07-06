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

// função recursiva.
func fibonacci(position uint) uint {
	if (position <= 1) {
		return position // retorna a propria posição quando for menor ou igual a 1.
	}

	return fibonacci(position - 2) + fibonacci(position - 1) // chamando a função recursiva e realizando o calculo de fibonacci.
} // não é recomendado para muitas execuções, pois é muito lento e pode cair em estouro de piha.

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

	// chamando a função recursiva.
	position1 := uint(10)

	for i := uint(0); i < position1; i++ {
		fmt.Println("O número fibonacci de", i, "é:", fibonacci(i)) // vai imprimir o número da sequência de fibonacci de 0 a 10.
	}

		fmt.Println("O número fibonacci de 20 é:", fibonacci(position1)) // 55
}