package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {

	var c float64 = float64(a) // TIPO(VARIÁVEL)
	fmt.Printf("O valor de C é: %g(%T) e o valor de B é: %s(%T)", c, c, b, b)

}
