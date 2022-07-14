package main

import "fmt"

func main() {
	var sueldo float32
	fmt.Println("Ingrese el sueldo:")
	fmt.Scan(&sueldo)
	if sueldo > 3000 {
		fmt.Print("Esta persona debe abonar impuestos")
	}
}
