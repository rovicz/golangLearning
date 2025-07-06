package main

import (
	"fmt"
	"reflect"
)

func main() {
	// padrão de criação e definição de uma variável/constante baseado em array.
	var array1 [5]int // [quantidade de elementos]tipo do elemento

	// maneira padrão de definição dos elementos de uma array.
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	fmt.Println(array1)

	// maneira mais prática de criação de uma variável e definição de elementos.
	array2 := [5]string{"um", "dois", "tres", "quatro", "cinco"} // [quantidade de elementos]tipo do elemento{elementos definidos}

	fmt.Println(array2)

	// utilizando o [...] você declara que a quantidade de elementos é igual a quantiade de elementos definidos dentro dessa array.
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array3)

	// slice NÃO é uma array, ele aponta para uma array, sendo uma referência. Possui tamanho dinâmico, diferente de uma array.

	// com o slice você não limita os dados, ele é bem mais utilizável e flexível que a array.
	slice1 := []string{"um", "dois", "tres", "quatro", "cinco", "seis", "sete"} // []tipo do elemento{elementos definidos}
	slice1 = append(slice1, "oito", "nove", "dez") // adicionando elementos ao slice via append.

	fmt.Println(slice1)

	fmt.Println( "Tipagem de uma Array:", reflect.TypeOf(array1)) // a tipagem de uma array.
	fmt.Println("Tipagem de um Slice:", reflect.TypeOf(slice1)) // a tipagem de uma slice.
}