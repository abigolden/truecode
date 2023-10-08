package main

import (
	"fmt"
	"practica3/practica3"
)

func main() {
	fmt.Println(ohm(0, 2, 3))
	fmt.Println(ohm(6, 2, 0))
	fmt.Println(ohm(6, 0, 3))
	fmt.Println(ohm(1, 2, 3))
	fmt.Println(ohm(0, 0, 3))

	///
	foo := practica3.Foobar(3)
	bar := practica3.Foobar(5)
	deLosDos := practica3.Foobar(15)
	ninguno := practica3.Foobar(11)

	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(deLosDos)
	fmt.Println(ninguno)

}

func ohm(v, r, i float64) float64 {
	if v == 0 {
		return i * r
	} else if i == 0 {
		return v / r
	} else if r == 0 {
		return v / i
	} else {
		return 0.0
	}
}
