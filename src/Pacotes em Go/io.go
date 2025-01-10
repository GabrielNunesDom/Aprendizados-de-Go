//O pacote Io fornece interface básicas primitivas
//WriteString escreve o conteúdo da string.
//

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}

}
