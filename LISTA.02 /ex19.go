package main

import f "fmt"
import m "math" 

func main() {
 var figura int
 var area, volume, raio, altura float64

    f.Print("Qual o tipo da figura? (1-cone / 2-cilindro / 3-esfera): ")
    f.Scan(&figura)

    switch figura {
    case 1:
        f.Print("Qual o valor do raio e da altura do cone? ")
        f.Scan(&raio, &altura)
        
        volume = (m.Pi * raio * raio * altura) / 3
        area = m.Pi*raio*(m.Sqrt((raio*raio)+(altura*altura))) + m.Pi*raio*raio

    case 2:
        f.Print("Qual o valor do raio e da altura do cilindro? ")
        f.Scan(&raio, &altura)
        
        volume = m.Pi * raio * raio * altura
        area = 2*m.Pi*raio*raio + 2*m.Pi*raio*altura

    case 3:
        f.Print("Qual o valor do raio da esfera? ")
        f.Scan(&raio)
        
        volume = (4 * m.Pi * raio * raio * raio) / 3
        area = 4 * m.Pi * raio * raio

    default:
        f.Println("Opção inválida!")
        return 
    }

    f.Printf("Area = %.2f \nVolume = %.2f\n", area, volume)
}
