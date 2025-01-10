// defer: escalona nossas funções
// função no dia-a-dia: fechar obrigatoriamente um arquivo depois que ele já foi aberto
package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia2()
	dia1()
}
