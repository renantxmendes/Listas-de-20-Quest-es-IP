package main

import "fmt"

var nota1, nota2, nota3, soma float64

func main() {
	fmt.Print("Escreva suas notas")
	fmt.Scan(&nota1, &nota2, &nota3)

	soma=(nota1 + nota2 + nota3) / 3

	if soma >= 6 {
		fmt.Printf("MEDIA = %.2f\nAPROVADO", soma)
	} else if soma<6 {
		fmt.Printf("MEDIA = %.2f\nREPROVADO", soma)
	}

}
