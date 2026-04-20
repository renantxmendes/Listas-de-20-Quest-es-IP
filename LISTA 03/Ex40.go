package main

import f "fmt"

func main() {
	preco := 6.00
	ingressos := 130
	despesas := 300.00

	var lucroMaximo, precoMaximo float64
	var ingressosMaximo int
	primeiraIteracao := true

	f.Println("Preco (R$) | Ingressos | Lucro Esperado (R$)")
	for preco >= 1.00 {
		receita := float64(ingressos) * preco
		lucro := receita - despesas

		f.Printf("%-10.2f | %-9d | %.2f\n", preco, ingressos, lucro)

		if primeiraIteracao || lucro > lucroMaximo {
			lucroMaximo = lucro
			precoMaximo = preco
			ingressosMaximo = ingressos
			primeiraIteracao = false
		}

		preco = preco - 0.60
		ingressos = ingressos + 30
	}
	f.Printf("\n--- Melhor Cenario ---\n")
	f.Printf("Lucro Maximo: R$ %.2f\n", lucroMaximo)
	f.Printf("Preco Ideal: R$ %.2f\n", precoMaximo)
	f.Printf("Ingressos Vendidos: %d\n", ingressosMaximo)
}
