package main

import "fmt"

func main() {
	var s []int
	printAppendedSlice(s)

	s = append(s, 0)
	printAppendedSlice(s)

	s = append(s, 2, 3, 4)
	printAppendedSlice(s)
}

func printAppendedSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}