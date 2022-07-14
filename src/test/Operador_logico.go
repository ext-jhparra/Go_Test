package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&num2)
	fmt.Print("Ingrese tercer valor:")
	fmt.Scan(&num3)
	fmt.Print("El mayor de los tres es:")
	if num1 > num2 && num1 > num3 {
		fmt.Print(num1)
	} else {
		if num2 > num3 {
			fmt.Print(num2)
		} else {
			fmt.Print(num3)
		}
	}

	var dia, mes, año int
	fmt.Print("Ingrese número de día:")
	fmt.Scan(&dia)
	fmt.Print("Ingrese número de mes:")
	fmt.Scan(&mes)
	fmt.Print("Ingrese número de año:")
	fmt.Scan(&año)
	if mes == 1 || mes == 2 || mes == 3 {
		fmt.Print("Corresponde al primer trimestre")
	}
}
