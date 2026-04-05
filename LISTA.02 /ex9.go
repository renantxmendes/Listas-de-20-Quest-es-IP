func main() {
    var val, lucro float64
    f.Print("Digite o valor de sua compra: ")
    f.Scan(&val)
    
    if val < 10 {
        lucro = val * 1.7
    } else if val >= 10 && val < 30 {
        lucro = val * 1.5
    } else if val >= 30 && val < 50 {
        lucro = val * 1.4
    } else {
        lucro = val * 1.3
    }
    f.Printf("O lucro será de: %.2f\n", lucro)
}
