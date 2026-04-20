package main

import f "fmt"

func main() {
	var totalGraos uint64
	var graosQuadroAtual uint64 = 1

	for i := 1; i <= 64; i++ {
		totalGraos = totalGraos + graosQuadroAtual
		graosQuadroAtual = graosQuadroAtual * 2
	}

	f.Printf("Total de graos de milho no tabuleiro: %v\n", totalGraos)
}
