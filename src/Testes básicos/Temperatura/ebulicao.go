// declarar meu pacote principal
package main

//importar a função fmt
import "fmt"

//declarar da variável CONST do ponto de ebulição da água em  F
const ebulicaoF = 212.0

//função principal
func main() {

	tempF := ebulicaoF            //variável do valor da temperatura em graus F
	tempC := (tempF - 32) * 5 / 9 //Conversão de F para C
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em F° é = %g A temperatura de ebulição da água em C° é = %g", tempF, tempC)
	//A gente espera que apareça F = 212 e temperatura em C = 100

}
