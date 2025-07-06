package main

import "fmt"

func main() {
	// declarei uma variável de tipo int e atribui um valor.
	var var1 int = 10
	var var2 int = var1 // devlarei uma variável de mesmo tipo e atribui o valor da outra.

	fmt.Println(var1, var2) // logo as duas variáveis tem o mesmo valor..
	
	var1++ // aumentei o valor da variável var1 em 1.
	fmt.Println(var1, var2) // aumentou o valor da variável var1, mas não do var2.
	// por lógica, ele está atribuindo o valor inicial de var1 como var2, logo qualquer alteração não afeta a outra.

	// os PONTEIROS são uma REFERÊNCIA de memória, não o valor.
	var var3 int
	var ponteiro *int // o * declara como ponteiro.

	var3 = 1000 // atribui o valor 10 a var3.
	ponteiro = &var3 // atribui o endereço da variável var3 a ponteiro utilizando o "E comercial".

	fmt.Println(var3, ponteiro) // o ponteiro retorna o endereço de onde o valor de var3 está armazenado.
	fmt.Println(*ponteiro) // utilizando o *, ele retorna o valor do ponteiro, realizando uma desreferenciação.

	// entendendo isso, o ponteiro aponta para o LOCAL onde o valor da variável está armazenado, não o valor, sendo assim, qualquer alterção do valor
	// de uma variável, vai ser acompanhada pela extração via desreferenciação do local onde está armazenado.
}