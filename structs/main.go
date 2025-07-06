package main

import "fmt"

// declarando um type (no caso no go como struct) para um modelo simples de usuário contendo name e email.
type user struct {
	name string
	mail string
} // pode criar struct dentro de struct, seguindo a mesma ideia de um type/interface do typescript por exemplo.

func main() {
	// declarando uma variável de tipo user de forma básica.
	var victor user
	victor.name = "Victor"
	victor.mail = "victor@gmail.com"

	fmt.Println("Nome do Usuário 1:", victor.name)
	fmt.Println("Email do Usuário 1:", victor.mail)

	// declarando uma variável de tipo user de forma mais agilidosa e prática.
	julia := user{"Júlia", "juju@gmail.com"}

	fmt.Println("Nome do Usuário 2:", julia.name)
	fmt.Println("Email do Usuário 2:", julia.mail)

	// declarando a variável com somente um campo sendo utilizado, caso você não tenha todos dados necessários do struct.
	chloe := user{name: "Chloe"} // utilize a propria propriedade do struct e passe o valor (prop: value).

	fmt.Println("Nome do Usuário 3:", chloe.name)
}