//Select funciona como um Switch para canais.
//Select permite que você guarde operações de vários canais.
//Combinar goroutines e canais com select é um recurso poderoso do GO.
//Selecionaremos dois canais para nosso exemplo

package main

import (
	"fmt"
	"time"
)

// Cada canal receberá um valor após algum tempo, para simular, por exemplo, o bloqueio de operações RPC em execução em goroutines concorrentes.
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //Recebemos os valores "um" e depois "dois" conforme o esperado.
		time.Sleep(1 * time.Second) //Observe que o tempo total de execução é 2 segundos, pois,
		//o 1 e o 2 segundos são Sleeps executados simultaneamente.
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select { //Usaremos select para aguardar esses dois valores simultaneamente, imprimindo cada um a medida que chegarem}
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}
}
