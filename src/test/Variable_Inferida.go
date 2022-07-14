package main

import "fmt"

func main() {
	var valor1, valor2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&valor2)
	suma := valor1 + valor2
	fmt.Println("La suma de los dos valores es:", suma)
	resta := valor1 - valor2
	fmt.Println("La diferencia del primer valor respecto al segundo es:", resta)
}
