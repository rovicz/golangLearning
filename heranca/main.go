package main

import "fmt"

// essa é a maneira mais próxima do que é herança dentro do go.
type person struct {
	name string
	age  uint8
}

// a fim de não tornar redundante e passar os mesmos types para o struct student, apenas passo o struct person e adiciono o que for necessário para esse novo struct.
type student struct {
	person
	school string
}

func main() {
	// criando uma variável com base na person.
	victorPerson := person{"Victor", 20}

	// criando uma varíavel com base na student, absorvendo a variável da person e complementando com que a struct de student pede.
	victor := student{victorPerson, "UNIFAN"}

	// de forma mais prática e ágil, pode-se aplicar assim.
	julia := student{person{"Júlia", 20}, "UNIFAN"}

	fmt.Println(victor)
	fmt.Println(julia)
}