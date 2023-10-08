package main

import (
	"fmt"
)

func main() {
	var altura int
	fmt.Println("Enter the height of the Mario's pyramid (min.1, max.8)")
	fmt.Scan(&altura)

	if altura < 1 || altura > 8 {
		fmt.Println("The height should be between 1 y 8.")
		return
	}

	for i := 1; i <= altura; i++ {
		// esta parte imprimo los espacios para que este alineado a la derecha
		for j := 1; j <= altura-i; j++ {
			fmt.Print(" ")
		}

		// aca imprimo los #
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
