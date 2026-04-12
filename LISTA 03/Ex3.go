package main

import f "fmt"

func main() {

	var sal_c, sal_j float64
	var meses int 

    f.Print("Digite o salário de Carlos: ")
	f.Scan(&sal_c)

	sal_j = sal_c /3 

	for sal_j < sal_c {
	
	 sal_j = sal_j * 1.05
	 sal_c = sal_c * 1.02
	 meses++
    
	}
    f.Print(" Em ", meses," meses o salário de João será igual ou maior o de Carlos. ")
	
}
