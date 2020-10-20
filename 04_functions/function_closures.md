# Function closures
A function value that references variables from outside its body

Functions are values:
```go
func compute(fn func(x int, y int) int) int {
    return fn(1, 1)
}
```