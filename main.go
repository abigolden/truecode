package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Escribe el nombre del proyecto a ejecutar")

		return
	}

	projectName := os.Args[1]

	switch projectName {
	case "bmi":
		fmt.Println("Ejecutar bmi")
	case "evenodd":
		fmt.Println("Ejecutar evenood")
	case "foobar":
		fmt.Println("Ejecutar foobar")
	case "mario":
		fmt.Println("Ejecutar maro")
	case "ohm":
		fmt.Println("Ejecutar ohm")
	default:
		fmt.Println("No existe el proyecto")
	}
}
