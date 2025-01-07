package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {

	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(pessoa{"João", 20})

}

//Estrutura da estrura: type(declarar uma variavel) struct
//Estruturas são coleções de "campos"
//São utéis para: agrupar dados, formar registros
