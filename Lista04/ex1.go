package main
import f "fmt"
func calculaPotencia(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * calculaPotencia(x, n-1)
}
func main() {
	var x, n int

	f.Print("Digite a base (x): ")
	f.Scan(&x)
	f.Print("Digite o expoente (n): ")
	f.Scan(&n)

	resultado := calculaPotencia(x, n)

	f.Printf("Resultado: %d\n", resultado)
}
