package main

import "fmt"

func main() {

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 1)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)

}

// Como representar a fatia: var x []float64
// fatia:=make([]float64, tamanho da fatia do array)
// fatia:=[low:high]
// fatia:=arr[0:5]
//acrescentar elemento | Usar a palavra reservada: APPEND
//copiar uma fatia para a outra COPY
