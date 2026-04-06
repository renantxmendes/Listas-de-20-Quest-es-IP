package main
import f "fmt"
func main() {
    var valor_prod float64
    var met_pag string

    f.Print("Digite o valor do produto: ")
    f.Scan(&valor_prod)
    
    f.Print("Forma de pagamento (Dinheiro, Cheque, Crédito, 2x, 3x): ")
    f.Scan(&met_pag)

    switch met_pag {
    case "Dinheiro", "Cheque":
        f.Printf("Preço com 10%% de desconto: %.2f\n", valor_prod*0.9)
        
    case "Crédito":
        f.Printf("Preço com 5%% de desconto: %.2f\n", valor_prod*0.95)
        
    case "2x": 
        f.Printf("Duas parcelas de: %.2f\n", valor_prod/2)
        
    case "3x":
        val_juros := valor_prod * 1.1
        f.Printf("Três parcelas de: %.2f\n", val_juros/3)
        
    default:
        f.Println("Forma de pagamento inválida!")
    }
}
