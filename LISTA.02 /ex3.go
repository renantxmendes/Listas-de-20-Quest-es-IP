package main
import f "fmt"

func main() {
	var a , b , soma int 
	f.Print("Digite um número inteiro qualquer: ")
	f.Scan(&a)
	f.Print("Digite outro número inteiro: ")
	f.Scan(&b)
	soma = a + b
	if soma > 20{
		soma = soma + 8
		f.Printf("O resultado será de: %d ",soma)
    }else{
		if soma <= 20{
			soma = soma - 5
		f.Printf("O resultado será de: %d  ",soma)
		}
	}
	

	
	
}
