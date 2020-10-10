package main

import "fmt"

func main() {
	// Arrays
	var fruit [2] string

	// Assigns values
	fruit[0] = "Apple"
	fruit[1] = "Peer"

	fmt.Println(fruit)

	// Declare and assign
	vegetable := [2]string{"Potato", "Carrot"}

	fmt.Println(vegetable)


	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2]) // Starts at one and stops before 2
	fmt.Println(fruitSlice[1:3]) // Starts at one and stops before 3
}
