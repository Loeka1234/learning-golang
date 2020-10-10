package main

import "fmt"

func main() {
	// Using var or const
	var name = "Loeka"
	var nr int = 2
	const isCool = true

	// Shorthand
	age := 17
	hello := "World"
	foo, bar := "foo", "bar"

	fmt.Println(name, age, nr)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)
	fmt.Println(hello, foo, bar)
}
