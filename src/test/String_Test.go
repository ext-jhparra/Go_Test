package main

import "fmt"

func main() {
	var nombre1, nombre2 string
	fmt.Print("Ingrese primer nombre:")
	fmt.Scan(&nombre1)
	fmt.Print("Ingrese segundo nombre:")
	fmt.Scan(&nombre2)
	if nombre1 < nombre2 {
		fmt.Print(nombre1, " es menor alfabéticamente hablando")
	} else {
		fmt.Print(nombre2, " es menor alfabéticamente hablando.")
	}
}
