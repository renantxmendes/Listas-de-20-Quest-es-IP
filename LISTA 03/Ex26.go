package main

import f "fmt"

func main() {
	soma := 0.0
	num := 100.0
	fat := 1.0

	f.Println("Calculando os 20 primeiros termos")

	for i := 1; i <= 20; i++{
		termoAtual := num / fat
		soma = soma + termoAtual

		f.Printf("Termo %2d: %3.0f / %-18.0f \n", i, num, fat)

		num = num - 1.0 
		
		fat = fat * float64(i) 
	}
	f.Printf("\nA soma dos 20 primeiros termos é: %f\n", soma)
}
