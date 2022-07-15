package main

import "fmt"

func main() {
	var mat [3][5]int
	for f := 0; f < 3; f++ {
		for c := 0; c < 5; c++ {
			fmt.Print("Ingrese componente:")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Println("Impresión de la matriz completa con la función Println")
	fmt.Println(mat)
	fmt.Println("Impresión elemento a elemento")
	for f := 0; f < 3; f++ {
		for c := 0; c < 5; c++ {
			fmt.Print(mat[f][c], " ")
		}
		fmt.Println()
	}
}
