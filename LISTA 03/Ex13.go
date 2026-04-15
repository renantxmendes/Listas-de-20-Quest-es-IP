package main
import f "fmt"
func main() {
var H, numerador float64 
numerador = 1.0

f.Print("Calculando a série H...")

for denominador:=1; denominador <= 50; denominador++{
	termo:= numerador/float64(denominador)
	H += termo 
	numerador += 2.0

}
f.Printf("O valor final da série H é: %.4f\n",H)
