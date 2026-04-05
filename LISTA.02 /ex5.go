package main
import f "fmt"
func main() {
	var num int
	f.Print("Escreva um número inteiro: ")
	f.Scan(&num)
	if num % 5 ==  0 {
		f.Printf("O número é divisível por 5")
	}else{
		f.Print("O número não é divisível por 5")
	}
	
}
