package main

import (
	"fmt"
	"time"
)

func pingar(c1, c2 chan string) {
	for i := 0; i < 2; i++ {
		c1 <- "ping"
		c2 <- "pong"
	}
}
func imprimir(c1, c2 chan string) {
	for {
		msg1 := <-c1
		fmt.Println(msg1)
		time.Sleep(time.Second * 1)
		msg2 := <-c2
		fmt.Println(msg2)
		time.Sleep(time.Second * 2)
	}
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go pingar(c1, c2)
	go imprimir(c1, c2)

	var entrada string
	fmt.Scanln(&entrada)
}
