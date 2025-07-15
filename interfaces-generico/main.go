package main

import "fmt"

// as fnções genericas são funções que podem ser utilizadas com qualquer tipo de dados.
// pontualmente utilizadas, mas nem sempre recomendado, pois podem ser mais difíceis de entender e bagunçadas, como gambiarras.

func generic(interf interface{}) { // func nomeFunc(nome do parametro interface{}) {...}.
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(123)
	generic(true)

	// o fmt.Println recebe um argumento do tipo interface{}, que pode ser qualquer tipo de dado, assim imprimindo QUALQUER dado que ele receber.
	fmt.Println(123, "String", true)
}