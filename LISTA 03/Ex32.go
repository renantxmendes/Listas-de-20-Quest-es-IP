package main

import f "fmt"

func main() {
	var n1, n2 int

	f.Print("Digite o primeiro numero (N1): ")
	f.Scan(&n1)
	f.Print("Digite o segundo numero (N2): ")
	f.Scan(&n2)

	resultado := 0

	for i := 0; i < n2; i++ {
		resultado = resultado + n1
	}

	f.Printf("Resultado da multiplicacao: %d\n", resultado)
}
