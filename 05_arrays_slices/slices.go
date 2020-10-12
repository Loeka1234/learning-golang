package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A slice
	// Dynamically-sized
	// Much more common than arrays
	var mySlice []int = primes[1:4] // Includes index 1 but exluces index 4
	fmt.Print(mySlice)
}
