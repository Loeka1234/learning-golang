package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceWithName("a", a)

	b := make([]int, 0, 5)
	printSliceWithName("b", b)

	c := b[:2]
	printSliceWithName("c", c)

	d := c[2:5]
	printSliceWithName("d", d)
}

func printSliceWithName(s string, x []int) {
	fmt.Printf("%s len=%d, cap=%d %v\n", s, len(x), cap(x), x)
}