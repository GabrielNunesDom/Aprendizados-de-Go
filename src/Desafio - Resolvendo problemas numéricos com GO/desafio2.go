package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPam", i)
		} else if i%3 == 0 {
			fmt.Println("Pin", i)
		} else if i%5 == 0 {
			fmt.Println("Pam", i)
		}
	}
}
