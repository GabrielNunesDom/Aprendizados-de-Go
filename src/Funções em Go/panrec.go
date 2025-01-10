//panic: erro do programador/erro de tempo de execução ou execução
//recover: interrompe o panic

package main

import "fmt"

func main() {
	defer func() {

		x := recover()
		fmt.Println(x)

	}()
	panic("Pânico")
}
