//Interfaces são coleções de métodos

package main

import (
	"fmt"
	"math" //utilizado para calcular áreas de figuras geométricas. Usado com o Pi.
)

//Aqui temos uma interface usada para formas geométricas

type geometria interface {
	area() float64
	perimetro() float64
}

//Para essa aula vamos usar a interface nos tipos QUADRADO e CÍRCULO.
// Área de um quadrado: lado² ou lado*lado
//perímetro = soma dos lados

type quadrado struct {
	lado float64
}
type circulo struct { //área círculo: (Pi)*raio² perímetrodo círculo: 2*r*(pi)
	raio float64
}

//Para usar uma interface em GO, só precisamos usar todos os métodos da interface.
//Aqui nós usaremos `geometria` em `quadrados`

// A implementação do `quadrado`
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// A implementação do `círculo`
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

//Se a variável tem um tipo interface, então nós podemos chamar métodos que estão na interface nomeada.
//Aqui temos uma função genérica `medida` tomando vantagem desse trabalho em qualquer `geometria`.

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 25}
	c := circulo{raio: 100}
	//Os tipos de estrutura `círculo` e `quadrado` ambos implementam a interface `geometria`,
	//então, nós podemos usar instâncias dessas estruturas como argumentos para ``medir`.

	medir(q)
	medir(c)
}
