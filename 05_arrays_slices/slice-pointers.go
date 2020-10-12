package main

import "fmt"

// Slices do not store any data
// It describes a section of an underlying array
// Changing the elements of a slice modifies the corresponding elements of its underlying array
// Other slices that share the same underlying array will see those changes
func main() {
	names := [4]string{
		"Loeka",
		"Jb",
		"Casper",
		"Arthur",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
