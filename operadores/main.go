package main

import "fmt"

func main() {
	// arimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	restoDivisao := 1 % 2

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(restoDivisao)

	// não pode somar/subtrair/multiplicar/dividir variaveis de tipagens diferentes.
	var numero16 int16 = 10
	var numero32 int32 = 10
	//soma1632 = numero16 + numero32 (não pode).

	fmt.Println(numero16)
	fmt.Println(numero32)

	// atribuição
	// vc pode definir o valor tipado manualmente ou utilizar o ":=" para atribuir o valor do tipo inferido.
	var variavel string = "String"
	variavel2 := "String2"

	fmt.Println(variavel)
	fmt.Println(variavel2)

	// operadores relacionais
	fmt.Println(1 == 1)
	fmt.Println(1 != 1) // diferente do js/ts, o diferente é != em vez de !==
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)

	// operadores lógicos
	fmt.Println(true && false) // se as duas forem true, retorna true = false
	fmt.Println(true || false) // se uma forem true, retorna true = true
	fmt.Println(!true) // se for false, retorna true = false

	// operadores unários
	numeroUnario := 1

	numeroUnario++
	numeroUnario += 15 // numeroUnario = numeroUnario + 15

	numeroUnario--
	numeroUnario -= 5 // numeroUnario = numeroUnario - 5
	fmt.Println(numeroUnario) // funciona a qualquer calculo.

	// não existem operadores ternários em GO.
	// dito isso, use if/else.
	var textoVariavel string 
	if numeroUnario > 10 {
		textoVariavel = "maior que 10"
	} else {
		textoVariavel = "menor que 10"
	} // bem feio, mas é a simplicidade do Go de fazer tudo de uma maneira e prática só.

	fmt.Println(textoVariavel)
}