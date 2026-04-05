package main

import f "fmt"
var num int 

func main(){
f.Print("Digite um número inteiro qualquer: ")	
f.Scan(&num)
        if num == 0{
        f.Printf("Esse número é NULO")
}else{
	    if num > 0{
		f.Printf("Esse número é POSITIVO")
	}else{
		if num < 0{
		f.Printf("Esse número é NEGATIVO")	
		}
	}
}
	
	}
