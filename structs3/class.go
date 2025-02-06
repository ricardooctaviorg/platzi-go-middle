package main

import "math"

type Cuadrilatero struct {
	base   float64
	altura float64
}

type Cuadrado struct {
	Cuadrilatero
}

type Rectangulo struct {
	Cuadrilatero
}

type Calculate interface {
	CalculateArea() float64
}

func (cuadrado Cuadrado) CalculateArea() float64 {
	return math.Round(math.Pow(cuadrado.base, 2))
}

func (rectangulo Rectangulo) CalculateArea() float64 {
	return math.Round(rectangulo.base * rectangulo.altura)
}

func CalculateArea(calculate Calculate) {
	println(calculate.CalculateArea())
}

func newCuadrado(base float64) *Cuadrado {
	return &Cuadrado{Cuadrilatero{base: base}}
}

func newRectangulo(base float64, altura float64) *Rectangulo {
	return &Rectangulo{Cuadrilatero{base: base, altura: altura}}
}

func main() {
	cuadrado := newCuadrado(10)
	rectangulo := newRectangulo(100, 200)

	CalculateArea(cuadrado)
	CalculateArea(rectangulo)
}
