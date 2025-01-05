// declarando meu pacote principal
package main

//importando a função fmt
import "fmt"

//declarando a variável CONST do ponto de ebulição da água em  K
const ebulicaoK = 372.0

//função principal
func main() {

	tempK := ebulicaoK     //variável do valor da temperatura em graus K
	tempC := (tempK - 273) //Conversão de K para C
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em K° é = %g A temperatura de ebulição da água em C° é = %g", tempK, tempC)
	//A gente espera que apareça F = 212 e temperatura em C = 100

}
