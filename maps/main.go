package main

import "fmt"

func main() {
	user := map[string]string{  //map[tipo da chave]tipo do valor
		"name": "Victor",
		"city": "Goiânia",
	}

	fmt.Println("map completo do user:", user)
	fmt.Println("nome do user:", user["name"]) // não é possível acessar como por exemplo "user.name", deve-se acessar como user["name"].

	user2 := map[string]map[string]string { //map[tipo da chave]map[tipo da chave]tipo do valor
		"name": {
			"first": "Victor",
			"last": "Fernandes",
		},
	}

	fmt.Println("map de mapa do user2:", user2)
	fmt.Println("primeiro nome do user2:", user2["name"]["first"]) // procurando dentro de name o first.
	fmt.Println("ultimo nome do user2:", user2["name"]["last"]) // procurando dentro de name o last.

	// delete(user2, "name") // remove o "name" do map do user2..
}