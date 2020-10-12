# Slice literals
A slice literal is like an array literal without the length.
```go
[]bool{true, true, false}
```
Thus creates an array, then builds a slice that references it

# Slice defaults
```go
var a [10]int

// Slices expressions that are equivalent:
a[0:10]
a[:10]
a[0:]
a[:]
```

# Nil slices
The zero value of a slice is `nil`
A nil slice has a length and capacity of 0 and has no underlying array

# `make` slice
Slices can be created with the built-in `make` function
