package goarea

import "math"

const Pi = 3.141516

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect ao criar um comentario com o nome da função, é exibido quando posicionar o mouse em cima da mesma
func Rect(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
