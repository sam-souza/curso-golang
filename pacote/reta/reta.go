package main

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visivel dentro e fora do pacote)!
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)!

// Exemplo
// Ponto 			- público
// ponto ou _ponto 	- privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
