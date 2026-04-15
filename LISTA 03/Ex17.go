package main
import f "fmt"
func main() {
var n1, n2 int

f.Print("Digite um número inteiro: ")
f.Scan(&n1)
f.Print("Digite novamente outro número inteiro: ")
f.Scan(&n2)
// Swap
if n1 > n2{
	temp:= n1
	n1 = n2
	n2 = temp
}
f.Printf("\nNúmeros primos entre %d e %d:\n", n1, n2)

for numero:=n1; numero <=n2; numero++{
	if numero < 2{
		continue
	}
	ehPrimo:= true 	

	for divisor:= 2; divisor < numero; divisor++{
		if numero%divisor == 0{
			ehPrimo = false
			break
		}

	}
    if ehPrimo == true{
		f.Print(numero, "  ")
	}
  } 
}
