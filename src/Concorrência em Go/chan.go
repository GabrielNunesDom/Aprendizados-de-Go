//canal: modo de duas goroutines se comunicarem e sincronizarem sua execução.

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { // Palavra reservada para canal: chan
	for i := 0; ; i++ {
		c <- "ping" //usado para enviar e receber mensagem pelo canal
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
