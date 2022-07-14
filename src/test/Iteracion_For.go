package main

import "fmt"

func main() {

	for x := 1; x <= 100; x++ {
		if x == 100 {
			fmt.Print(x)
		} else {
			fmt.Print(x, "-")
		}
	}
}
