package main

import "fmt"

var destino, ida_volta, valor int


func main() {
	fmt.Print("Digite o destino e o tipo de viagem: ")
	fmt.Scan(&destino, &ida_volta)

	if destino == 0 || destino > 4 || ida_volta < 0 || ida_volta > 1 {
		fmt.Print("Valores inválidos")
		return
	}
	if destino == 1 {
		valor = 500 
		if ida_volta == 1 {
			valor = 900
		}
	} else 
	if destino == 2 {
		valor = 350
		if ida_volta == 1 {
			valor = 650
		}
	} else 
	if destino == 3 {
		valor = 350
		if ida_volta == 1 {
			valor = 600
		}
	} else 
	if destino == 4 {
		valor = 300
		if ida_volta == 1 {
			valor = 550
		}
	}

	fmt.Println("O valor da viagem é: ", valor)

}
