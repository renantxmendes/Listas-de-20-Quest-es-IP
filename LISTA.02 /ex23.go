package main
import f "fmt"
func main() {
	var idade int
	var clas_eleit string

	f.Print("Digite sua idade: ")
	f.Scan(&idade)

	if idade < 0 {
		f.Println("Idade inválida")
		return
	}

	if idade < 16 {
		clas_eleit = "Não-eleitor"
	} else if idade < 18 || idade >= 65 {
		clas_eleit = "Eleitor facultativo"
	} else {
		clas_eleit = "Eleitor obrigatório"
	}

	f.Printf("Classe eleitoral: %s\n", clas_eleit)
}
