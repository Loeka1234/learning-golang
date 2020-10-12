package main

import "fmt"

func main() {
	// Array with length of 2
	// Arrays cannot be resized
	var myArr [2]string
	myArr[0] = "Hello"
	myArr[1] = "World!"
	fmt.Println(myArr[0], myArr[1])
	fmt.Println(myArr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	for _, p := range primes {
		fmt.Println(p)
	}
}
