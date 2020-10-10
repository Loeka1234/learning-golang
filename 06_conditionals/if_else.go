package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {
	if value := math.Pow(x, n); value < limit {
		return value
	} else {
		fmt.Printf("%g >= %g\n", value, limit)
	}
	return limit
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
