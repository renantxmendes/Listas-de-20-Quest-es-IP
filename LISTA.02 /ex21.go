package main
import f "fmt"
func main() {
	var num_id int
	var nota1, nota2, nota3, med_ex, med_aprovado float64
	var conc, aprovacao string

	f.Print("Digite o número de identificação do aluno: ")
	f.Scan(&num_id)
	f.Print("Digite as 3 notas do aluno e a média dos exercícios: ")
	f.Scan(&nota1, &nota2, &nota3, &med_ex)

	medAprov = (nota1 + nota2*2 + nota3*3 + medEx) / 7

	if med_aprovado >= 9 {
		conc = "A"
	} else if med_aprovado >= 7.5 {
		conc = "B"
	} else if med_aprovado >= 6 {
		conc = "C"
	} else if med_aprovado >= 4 {
		conc = "D"
	} else {
		conc = "E"
	}

	switch conc {
	case "A", "B", "C":
		aprovacao = "APROVADO"
	case "D", "E":
		aprovacao = "REPROVADO"
	}

	f.Printf("Número do aluno: %d\n"+
		"Notas: %.2f, %.2f, %.2f\n"+
		"Média dos exercícios: %.2f\n"+
		"Média de aproveitamento: %.2f\n"+
		"Conceito: %s - %s\n",
		num_id, nota1, nota2, nota3, med_ex, med_aprovado, conc, aprovacao)
}
