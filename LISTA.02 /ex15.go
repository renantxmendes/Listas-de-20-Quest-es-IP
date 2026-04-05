package main

import f "fmt"

var dia, mes, ano int
var nome_mes string

func main() {
	f.Print("Informe a data, o mês e o ano: ")
	f.Scan(&dia, &mes, &ano)
	if dia == 0 || dia < 0 || dia > 31 {
		f.Print("Data inválida")
		return
	}
	switch mes {
	case 1:
		nome_mes = "Janeiro"
	case 2:
		nome_mes = "Fevereiro"
	case 3:
		nome_mes = "Março"
	case 4:
		nome_mes = "Abril"
	case 5:
		nome_mes = "Maio"
	case 6:
		nome_mes = "Junho"
	case 7:
		nome_mes = "Julho"
	case 8:
		nome_mes = "Agosto"
	case 9:
		nome_mes = "Setembro"
	case 10:
		nome_mes = "Outubro"
	case 11:
		nome_mes = "Novembro"
	case 12:
		nome_mes = "Dezembro"
	}
	f.Printf("%d de %s de %d", dia, nome_mes, ano)
}
