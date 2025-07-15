package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64 // um método area dentro da interface que posteriormente será implementado por cada struct.
}

func writeArea(s shape) {
	fmt.Printf("A área da forma é %0.2f.\n", s.area())
} 

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 { // método area correspondente ao struct rectangle, que calcula a área de um retângulo.
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 { // método area correspondente ao struct circle, que calcula a área de um círculo.
	return math.Pi * (c.radius * c.radius)
}

func main() {
	rectangle1 := rectangle{height: 10, width: 5}
	circle1 := circle{radius: 5}

	writeArea(rectangle1) // necessariamente para utilizar a função writeArea, o struct rectangle precisa possuir o método area().
	writeArea(circle1) // necessariamente para utilizar a função writeArea, o struct circle precisa possuir o método area().
}