package main

import (
	"fmt"
	"math"
)

type MyI interface {
	M()
}

type MyT struct {
	S string
}

func (t *MyT) M() {
	fmt.Println(t.S)
}

type MyF float64

func (f MyF) M() {
	fmt.Println(f)
}

func main() {
	var i MyI
	i = &MyT{"hello"}
	describe(i)
	i.M()

	i = MyF(math.Pi)
	describe(i)
	i.M()
}

func describe(i MyI) {
	fmt.Printf("(%v, %T)\n", i, i)
}
