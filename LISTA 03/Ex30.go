package main

import f "fmt"
import m "math"

func main() {
	f.Println("Raio (cm) | Volume (cm3)")
	f.Println("------------------------")

	for r := 0.0; r <= 20.0; r = r + 0.5 {
		volume := (4.0 / 3.0) * m.Pi * m.Pow(r, 3.0)
		f.Printf("%-9.1f | %.2f\n", r, volume)
	}
}
