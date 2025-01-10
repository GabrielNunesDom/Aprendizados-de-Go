//implementa rotinas de utilitário para manipular caminhos de nome de arquivo de forma compatível com os caminhos de arquivo definidos pelo sistema operacional de destino
//O pacote de caminho de arquivo usa barras ou barras invertidas, dependendo do sistema operacional.
//Para processar caminhos com URLs que sempre usam barras independentemente do sistema operacional, consulte o pacote de caminhos.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
