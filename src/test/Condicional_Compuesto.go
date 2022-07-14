package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Print("El mayor es ", num1)
	} else {
		fmt.Print("El manor es ", num2)
	}
}
