package main

import "fmt"

//func [nombre de la función] ([parámetros de la función]) [valor que retorna] {
//[altoritmo]
//}
func retornarMayor(v1, v2 int) int {
	var mayor int
	if v1 > v2 {
		mayor = v1
	} else {
		mayor = v2
	}
	return mayor
}

func main() {
	var valor1, valor2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&valor2)
	fmt.Print("El valor mayor es ", retornarMayor(valor1, valor2))
}
