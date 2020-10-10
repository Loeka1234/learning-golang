package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less then %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d\n", y, x)
	}

	// Switch statement
	color := "red"
	switch color {
	case "red":
		fmt.Println("Color is red!")
		break
	case "green":
		fmt.Println("Color is green")
		break
	case "blue":
		fmt.Println("Color is blue")
		break
	default:
		fmt.Println("Color is not red, green or blue")
	}
}
