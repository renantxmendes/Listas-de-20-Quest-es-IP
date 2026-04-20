package main
import f "fmt"
func main() {
	var n int

	f.Print("Digite um numero inteiro positivo na base 10: ")
	f.Scan(&n)

	if n == 0 {
		f.Println("Base 16: 0")
		return
	}

	hexadecimal := ""
	temp := n
	digitosHexa := "0123456789ABCDEF"

	for temp > 0 {
		resto := temp % 16
		hexadecimal = string(digitosHexa[resto]) + hexadecimal
		temp = temp / 16
	}
	f.Printf("Base 16: %s\n", hexadecimal)
}
