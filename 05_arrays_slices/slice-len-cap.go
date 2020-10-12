package main

import "fmt"

func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	printSlice(mySlice)

	// Slice the slice to give it zero length
	mySlice = mySlice[:0]
	printSlice(mySlice)

	// Extends it length
	mySlice = mySlice[:4]
	printSlice(mySlice)

	// Drop its first two values
	mySlice = mySlice[2:]
	printSlice(mySlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
