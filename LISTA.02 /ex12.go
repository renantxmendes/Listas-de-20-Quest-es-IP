package main
import f "fmt"

func main() {
    var idade int
    var categoria string 

    f.Print("Quantos anos você tem: ")
    f.Scan(&idade)

    if idade >= 0 && idade <= 2 {
        categoria = "Recém-nascido"
    } else if idade >= 3 && idade <= 11 {
        categoria = "Criança"
    } else if idade >= 12 && idade <= 19 {
        categoria = "Adolescente"
    } else if idade >= 20 && idade <= 55 {
        categoria = "Adulto"
    } else {
        categoria = "Idoso"
    }
    f.Printf("A sua faixa etária é: %s\n", categoria)
}
