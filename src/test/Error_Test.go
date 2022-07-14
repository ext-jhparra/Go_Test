package main

import "fmt"

func main() {
	var lado int
	var superficie int

	fmt.Println("Ingrese el valor del lado del cuadro :")
	fmt.Scan(&lado)
	superficie = lado * lado
	fmt.Println("La superficie del cuadrado es : ", superficie)
}
