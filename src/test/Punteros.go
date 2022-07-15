package main

import "fmt"

func main1() {
	valor1 := 10
	valor2 := 20
	var pe *int
	pe = &valor1
	fmt.Println("Lo apuntado por pe es:", &pe, " - ", *pe, " - ", pe)
	fmt.Println("La direccion que almacena pe es:", pe)
	pe = &valor2
	fmt.Println("Lo apuntado por pe es:", *pe)
	fmt.Println("La direccion que almacena pe es:", pe)

}

func main2() {
	var s1 string = "uno"
	var s2 string = "dos"
	var ps *string
	ps = &s1
	fmt.Println(s1)
	*ps = "tres"
	fmt.Println(s1)
	s1 = "cuatro"
	fmt.Println(*ps)
	ps = &s2
	fmt.Println(*ps)
}

func main() {
	var f int
	var pe *int
	pe = &f
	for *pe = 1; *pe <= 10; *pe = *pe + 1 {
		fmt.Println(f)
	}
}
