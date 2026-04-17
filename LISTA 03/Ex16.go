package main
import f "fmt"
func main() {
	var a1, a2, n int
  
  f.Print("Digite um número inteiro: ")
  f.Scan(&a1)
  
  f.Print("Digite um número inteiro: ")
  f.Scan(&a2)
  
  f.Print("Digite a quantidade de termos da sequência: ")
  f.Scan(&n)

  f.Print("Série de Fetuccine:  ", a1, a2)
  
  prev2 := a1
  prev1 := a2
   
 for i:=3; i<=n; i++{
  var termo_atual int
  if i%2==0{
    termo_atual = prev1 - prev2
  }else{
    termo_atual = prev1 + prev2
  }
  f.Println(" ", termo_atual)
  prev2 = prev1
  prev1= termo_atual
    
 }
	f.Println()
}
