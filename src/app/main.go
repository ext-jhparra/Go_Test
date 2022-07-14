package main

import "fmt"

func main() {

	var horasTrabajadas int
	var costoHora float32
	var sueldo float32

	fmt.Println("Ingrese las horas trabajadas por el empleado")
	fmt.Scan(&horasTrabajadas)
	fmt.Println("Ingrese el valor por hora trabajada")
	fmt.Scan(&costoHora)
	sueldo = costoHora * float32(horasTrabajadas)
	fmt.Println("el sueldo a pagar al empleado es de : ", sueldo)
}
