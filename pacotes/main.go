package main

import (
	"fmt"
	"modules/components/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Message bellow:")
	auxiliar.Write()

	// testando um pacote de validação de e-mails.
	erro := checkmail.ValidateFormat("contato@victorferreira.dev") // nil (null no golang)
	erro2 := checkmail.ValidateFormat("1234567890") // invalid format
	fmt.Println(erro, erro2)
}