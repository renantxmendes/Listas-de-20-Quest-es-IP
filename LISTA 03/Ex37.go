package main
import f "fmt"
import m "math"
func main() {
	var n int

	f.Print("Digite um numero inteiro positivo na base 8: ")
	f.Scan(&n)

	decimal := 0.0
	temp := n
	expoente := 0.0

	for temp > 0 {
		digito := temp % 10
		decimal = decimal + (float64(digito) * m.Pow(8.0, expoente))
		temp = temp / 10
		expoente++
	}
	f.Printf("Base 10: %.0f\n", decimal)
}
