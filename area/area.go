package area

import "math"

// Pi é uma proporção numerica definida pela relação entre o perimetro de uma circunferencia e seu diametro
const Pi = 3.1415

// Circ responsavel por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
