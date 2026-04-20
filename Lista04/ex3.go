package main
import f "fmt"

func inverteArray(arr []int, inicio, fim int) {
	if inicio >= fim {
		return
	}

	arr[inicio], arr[fim] = arr[fim], arr[inicio]

	inverteArray(arr, inicio+1, fim-1)
}

func main() {
	valores := []int{10, 20, 30, 40, 50}
	tamanho := len(valores)

	f.Printf("Array original: %v\n", valores)

	inverteArray(valores, 0, tamanho-1)

	f.Printf("Array invertido: %v\n", valores)
}
