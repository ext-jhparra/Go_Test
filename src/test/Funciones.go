package main

import "fmt"

//func [nombre de la función] () {
//[altoritmo]
//}

func presentacion() {
	fmt.Println("Programa que permite cargar dos valores por teclado.")
	fmt.Println("Efectua la suma de dos valores")
	fmt.Println("y muestra el resultado de la misma")
	fmt.Println("*******************************")
}

func cargaSuma() {
	var valor1, valor2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&valor2)
	suma := valor1 + valor2
	fmt.Println("La suma de los dos valores es:", suma)
}

func finalizacion() {
	fmt.Println("*******************************")
	fmt.Println("Gracias por utilizar este programa")
}

func main() {
	presentacion()
	cargaSuma()
	finalizacion()
}
