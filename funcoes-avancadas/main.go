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

// defer

func studientIsOkay(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O Resultado será retornado em breve.") // o defer atrasa a exebição da mensagem e só retorna ANTES de qualquer return dentro da função.
	fmt.Println("Média sendo calculada...")

	mid := (n1 + n2) / 2

	if mid >= 6 {
		return true // retorna a função do defer e após isso retorna o valor, no caso "true".
	}
	return false // retorna a função do defer e após isso retorna o valor, no caso "false".

	// então o retorno seria assim:

	// Média sendo calculada...
	// Média calculada. O Resultado será retornado em breve.
	// true or false.
}

// panic e recover

func tryingToRecover() {
	if r := recover(); r != nil { // gera um recover via if init, caso o recover seja diferente de null, ele salva a aplicação.
		defer fmt.Println("A aplicação foi recuperada com sucesso!")
		fmt.Println("Recuperando a aplicação...")
	}
}

func studientIsOkayUsingPanic(n1, n2 float32) bool {
	defer tryingToRecover() // recupera a aplicação caso ocorra um panic.
	mid := (n1 + n2) / 2

	if mid > 6 {
		return true
	} else if mid < 6 {
		return false
	}

	panic("A média é exatamente 6.") // mata a execução da aplicação.
}


// closure
func closure() func() {
	text := "Dentro da função closure."

	func1 := func() {
		fmt.Println(text)
	}

	return func1
}

// ponteirs

func changeNumber(number *int) { // define um ponteiro para a variável number - não necessita de retorno.
	*number = *number * -1 // retorna o valor da variável number e multiplica-o por -1.
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

	// chamando a função recursiva.
	position1 := uint(10)

	for i := uint(0); i < position1; i++ {
		fmt.Println("O número fibonacci de", i, "é:", fibonacci(i)) // vai imprimir o número da sequência de fibonacci de 0 a 10.
	}

	fmt.Println("O número fibonacci de 20 é:", fibonacci(position1)) // 55

	// defer - como é dito pela própria função, ele adia a execução de una função e potencialmente realiza o que vem a seguir dela.
	fmt.Println(studientIsOkay(6,6)) // true

	// panic e recover
	fmt.Println(studientIsOkayUsingPanic(6,6)) // A média é exatamente 6. Caso não use um recover, a aplicação morre.
	fmt.Println("Após execução da função studientIsOkayUsingPanicOrRecover utilizando panic.")

	// closure
	text := "Dentro da função main."
	fmt.Println(text)

	func2 := closure()
	func2() // imprime "Dentro da função closure." e não "Dentro da função main." já que a clousure é uma função anônima.

	// pointers
	number := 15
	changeNumber(&number) // altera a variável number de 15 para -15 (var * -1), passando a variável como referência de memória.
	fmt.Println(number) // o valor da variável number agora é -15.

	// & = onde a variável se encontra na memória.
	// * = o valor da variável.
	// realizando assim uma alteração do valor da variável mesmo que fora da função.
}