package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipages de numeros no go.
	// int8 = 8 bits
	// int16 = 16 bits
	// int32 = 32 bits
	// int64 = 64 bits
	// int = usa como base os bits do seu sistema operacional.
	var numero16 int = 100 // como estou em um pc de 64 bits, ele seria um int64.

	//uint8 o conjunto de todos os inteiros de 8 bits sem sinal (0 a 255)
	//uint16 o conjunto de todos os inteiros de 16 bits sem sinal (0 a 65535)
	//uint32 o conjunto de todos os inteiros de 32 bits sem sinal (0 a 4294967295)
	//uint64 o conjunto de todos os inteiros de 64 bits sem sinal (0 a 18446744073709551615)
	var numeroUint uint = 100 // o uint é usado para representar valores não negativos.

	// alias - existem alguns alias para certas tipagens de dados.
	const numeroRune rune = 10000 // o rune é usado para representar o int32.
	const numeroByte byte = 100 // o byte é usado para representar o uint8.


	// para float, existe a opção do float32 e float64.
	const numeroFloat32 float32 = 5000200.22
	const numeroFloat64 float64 = 1129123123.333

	// tipagem de strings.
	const str string = "go é uma linguagem de programação de código aberto, desenvolvida pela Google."

	// não existe o char como em outras tipagens.
	char := "a" // ele vai retornar o número do caractere na tabela ASCII.

	// o valor booleano atua normalmente, entre false e true.
	var bol bool = true // o valor inicial do bol é false, caso ele não seja pré definido.

	// no Go, o erro é uma tipagem.
	var erro error
	var erroCustom error = errors.New("Erro customizado!") // para criar um erro, deve-se utilizar o pacote errors do Go.

	// TODA tipagem do Go tem um valor zero.
	// string = ""
	// int = 0
	// float = 0.0
	// bol = false
	// error = nil (null)

	fmt.Println(numero16, numeroUint, numeroRune, numeroByte, numeroFloat32, numeroFloat64, str, char, bol, erro, erroCustom)
}