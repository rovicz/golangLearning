package main

import "fmt"

func main() {
	number := 5

	if number > 4 {
		fmt.Println("O número é maior que 4.")
	} else {
		fmt.Println("Menor ou igual que 4.")
	}

	// if init (inicia uma variável dentro de um if, limitando a variável ao escopo do if).
	if otherNumber := 10; otherNumber > 5 {
		fmt.Println("O outro número é maior que 5.")
	} else {
		fmt.Println("O outro número éenor ou igual que 5.")
	} // sendo assim, a variável otherNumber é limitada ao escopo do if, não podendo ser acessada fora do if.

	// fmt.Println(otherNumber) - não é possível acessar a variável.

	if thirdNumber := 20; thirdNumber == 20 {
		fmt.Println("O terceiro número é igual a 20.")
	} else if thirdNumber > 20 { // pode ser utilizado o if else no golang.
		fmt.Println("O terceiro número é maior que 20.")
	} else {
		fmt.Println("O terceiro número é menor ou igual que 20.")
	}
}