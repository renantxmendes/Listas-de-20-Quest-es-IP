package main

import f "fmt"

func main() {
	var n1, n2 int

	f.Print("Digite o valor de N1 (positivo): ")
	f.Scan(&n1)
	f.Print("Digite o valor de N2 (positivo): ")
	f.Scan(&n2)

	quociente := 0
	resto := n1

	for resto >= n2 {
		resto = resto - n2
		quociente = quociente + 1
	}

	f.Printf("Quociente(%d,%d) = %d\n", n1, n2, quociente)
	f.Printf("Resto(%d,%d) = %d\n", n1, n2, resto)
}
