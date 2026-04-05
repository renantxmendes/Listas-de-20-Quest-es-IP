package main
import f "fmt"
func main() {
var x , resultado float64
f.Print("Digite um valor para x: ")
f.Scan(&x)
resultado = 8/ (2 - x)
f.Printf("O resultado será: %.2f", resultado)
	
}
