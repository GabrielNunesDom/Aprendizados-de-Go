//O pacote bytes implementa funções para a manipulação de fatias de bytes.
// Title trata s como bytes codificados em UTF-8 e retorna um cópia com todas as letras Unicode mapeadas.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
}
