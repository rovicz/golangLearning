package main

import "fmt"

func setDayOfWeek(dayNumber int) string {
	// o switch funciona como em qualquer linguagem.
	switch dayNumber {
	case 1: // seguindo o mesmo padrão do uso de case.
		return "Sunday"
		// fallthrough - o fallthrough é usado para pular o código de uma casa para a próxima casa (o retorno então seria "Monday").
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default: 
		// um retorno default, já que a função necessita de um retorno, dito inicialmente em string.
		return "Invalid day number!"
	}

	// o break não é necessário dentro do golang, ele já funciona nativamente.
}

func main() {
	// chama a função com uma variável, passando um número inteiro.
	randomDay := setDayOfWeek(3)
	invalidDayNumber := setDayOfWeek(10)

	fmt.Println(randomDay) // Tuesday (3).
	fmt.Println(invalidDayNumber) // Invalid day number! (10).
}