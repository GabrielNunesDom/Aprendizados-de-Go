package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("Esse são todos os múltiplos de 3 entre 0 e 100:", i)
		}

	}
}
