package main

import  "fmt"

var num int

func main() {
	fmt.Print("Escreva um número: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Print("O NÚMERO E DIVISÍVEL")
	} else {
		fmt.Print("O NÚMERO NÃO É DIVISÍVEL")
	}
}
