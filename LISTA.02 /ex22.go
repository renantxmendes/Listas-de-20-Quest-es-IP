package main
import f "fmt"

func main() {
	var mat_func int
	var qnt_hr_ex, sal_bruto float64
	var sal_min float64 = 788
	var val_hr_ex  float64 = 10

	f.Print("Matrícula do funcionário: ")
	f.Scan(&mat_func)
	f.Print("Quantidade de horas-extras trabalhadas: ")
	f.Scan(&qnt_hr_ex)

	var sal_hr_ex float64 = qnt_hr_ex * val_hr_ex
	sal_bruto = 3*sal_min + sal_hr_ex
	
	var sal_liq = sal_bruto * 0.68

	f.Printf("Salário Bruto: %.2f\nSalário Líquido: %.2f\n", sal_bruto, sal_liq)
}
