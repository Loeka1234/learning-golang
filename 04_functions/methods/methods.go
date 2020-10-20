package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// !! All methods on a given type should have either value or pointer receivers, but not a mixture of both

// Value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receiver
// The method can modify the value that its receiver points to
// Avoids copying the value on each method call (can be more efficient if the receiver is a large struct)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
