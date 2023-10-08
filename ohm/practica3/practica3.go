package practica3

func Foobar(num int) string {
	if num%3 == 0 {
		return "Foo"
	} else if num%5 == 0 {
		return "Bar"
	} else if num%3 == 0 && num%5 == 0 {
		return "FooBar"
	} else {
		return "No es mutiplo ni de 3 ni de 5"
	}
}
