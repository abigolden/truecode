package main

import (
	"fmt"
)

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Println(evenOrOdd(1))  // "odd"
	fmt.Println(evenOrOdd(22)) // "even"
}
