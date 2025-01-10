//Closure: capacidade de uma função "chamar" uma variável que está em outra função.

package main

import "fmt"

func main() {
	x := 0

	numero := func() int { //lembrar de declarar o que tem dentro da função
		x++ //incremento somar 1
		return x
	}
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}
