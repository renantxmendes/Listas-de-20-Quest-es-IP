package main
import f "fmt"
func main() {
	var n int

	f.Print("Digite um numero inteiro positivo na base 10: ")
	f.Scan(&n)

	if n == 0 {
		f.Println("Base 2: 0")
		return
	}
	binario := ""
	temp := n

	for temp > 0 {
		resto := temp % 2
		binario = f.Sprintf("%d%s", resto, binario)
		temp = temp / 2
	}

	f.Printf("Base 2: %s\n", binario)
}
