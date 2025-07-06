package main

import (
	"fmt"
	"time"
)

// no golang, a ÚNICA estrutura de repetição é o for.

func main() {
	i := 0

	// enquanto i for menor que 10 executa o código a seguir.
	for i < 4 {
		i += 2
		fmt.Println("Incrementando o contador do i:", i)
		time.Sleep(time.Second) // espera 1 segundo.
	}

	// assim como o if init, também temos o for init.
	for j := 0; j < 2; j++ {
		fmt.Println("Incrementando o contador do j:", j)
		time.Sleep(time.Second) // espera 1 segundo.
	}

	names := [3]string{"Victor", "Chloe", "Ju"}

	// utilizando o for com range para consumir uma array.
	for _, name := range names {
		fmt.Println("O nome da pessoa é:", name)
	}

	// caso quando o range absorve uma string, onde ele vai avaliar cada caractere.
	for _, letter := range "Victor" {
		fmt.Println("O primeiro caractere da string é:", string(letter)) // se não utilizar a função string() ele retorna o valor da letra na tabela asc.
	}
}