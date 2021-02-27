package goarea

import "math"

//Pi constante PI
const Pi = 3.1415

//Circulo = calcular area do circulo
func Circulo(valor float64) float64 {

	return math.Pow(valor, 2) * Pi

}

//Retangulo = calcular area
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Triangulo = calcular area (não é visivel fora do pacote por começar com "_")
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
