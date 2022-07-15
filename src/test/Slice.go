package main

import "fmt"

func main() {

	var cantidad int
	var suma float64
	fmt.Print("Cuantas facturas se procesarán:")
	fmt.Scan(&cantidad)
	facturas := make([]float64, cantidad)
	for f := 0; f < len(facturas); f++ {
		fmt.Print("Ingrese importe de la factura:")
		fmt.Scan(&facturas[f])
		suma = suma + facturas[f]
	}
	fmt.Println("Listado completo de facturas ingresadas")
	fmt.Println(facturas)
	fmt.Println("Importe total de todas las facturas:", suma)
}
