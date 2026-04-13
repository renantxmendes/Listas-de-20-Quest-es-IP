package main

import f "fmt"

func main() {
	
	var num, a int
	var eTriangular bool = false  
    a = 1
	
	f.Print("Digite um número inteiro qualquer: ")
	f.Scan(&num)
	
	for a * (a+1) * (a+2) <= num{
		
		if a * (a+1) * (a+2) == num{
		eTriangular = true 
		
		}
		a++
	}
	if eTriangular == true{
		f.Println("O número é triangular")
	}else{
		f.Println("O número não é triangular")
	}
}	
