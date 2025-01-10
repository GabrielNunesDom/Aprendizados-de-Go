//aceita um conjunto de dados e o reduz a um tamanho fixo menor.
//hashes são usadas em tudo em programação, principalmente para buscar Dados e Detectar
//em GO são divididas em CRIPTOGRAFADAS e NÃO CRIPTOGRAFADAS
//lista das NÃO CRIP: adler32. crc32, crc64
//nesse código vamos criar uma HASH e escrever nossos dados nele.

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//criar hash
	h := crc32.NewIEEE()
	//escrever nossos dados no hash
	h.Write([]byte("código com pacote hash"))
	// calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
