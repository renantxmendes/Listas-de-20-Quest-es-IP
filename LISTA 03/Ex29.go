package main

import f "fmt"

func main() {
	var n int
	
	f.Print("Digite um numero inteiro positivo N: ")
	f.Scan(&n)

	soma := 0

	for i := 1; i <= n; i++ {
		soma = soma + i
	}

	f.Printf("A soma de 1 ate %d e: %d\n", n, soma)
}
