package main

import (
	"fmt"
	"reflect"
)

func main() {
	// com o slice você não limita os dados, ele é bem mais utilizável e flexível que a array.
	slice1 := []string{"um", "dois", "tres", "quatro", "cinco", "seis", "sete"} // []tipo do elemento{elementos definidos}
	slice1 = append(slice1, "oito", "nove", "dez") // adicionando elementos ao slice via append.

	fmt.Println(slice1)
	fmt.Println("Tipagem de um Slice:", reflect.TypeOf(slice1)) // a tipagem de uma slice.

	// arrays internos
	slice2 := make([]float32, 14, 15) // make([]tipo do elemento, quantidade de elementos, capacidade)
	fmt.Println(slice2)
	
	slice2 = append(slice2, 0) // 15 de 15 (tamanho x capacidade)
	slice2 = append(slice2, 0) // 16 de 32 (tamanho x capacidade) - estourou o limite de capacidade, logo a capacidade será tamanho * 2.
	fmt.Println("Length do Slice:", len(slice2)) // length
	fmt.Println("Capacity do Slice:", cap(slice2)) // capacity

	// quando você define uma capacidade para o slice, ele irá alocar o espaço na memória, se a capacidade for maior que
	// o tamanho do slice, ele irá realocar o espaço com o dobro do tamanho.

	// array = lista com tamanho fixo.
	// slice = lista com tamanho dinâmico.
}