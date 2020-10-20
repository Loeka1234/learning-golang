package main

import "fmt"

// My weird Fibonacci closure
// It looks a bit like the better Fibonacci I found on the internet but there are certain optimizations
func fibonacci() func() int {
	prev := 0
	nr := 1
	return func() int {
		newNr := nr + prev
		defer func() {
			prev = nr
			nr = newNr
		}()
		return prev
	}
}

// The better Fibonacci function I found on the internet
func betterFibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	f2 := betterFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f2())
	}
}
