package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// func (referência do struct struct) nomeDoMetodo(args) {...}

func (u user) getName() { // criando um método "getName" para o struct user.
	sentence := fmt.Sprintf("O nome do usuário é %s.", u.name)
	
	fmt.Println(sentence)
}

func (u user) getAge() { // criando um método "getAge" para o struct user.
	sentence := fmt.Sprintf("O usuário %s tem %d anos.", u.name, u.age)

	fmt.Println(sentence)
}

func main() {
	user1 := user{"Victor", 20}

	user.getName(user1) // imprimiria "O nome do usuário é Victor."
	user.getAge(user1) // imprimiria "O usuário Victor tem 20 anos."
}