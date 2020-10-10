# Conditionals
### if
If with a short statement:
```go
func pow(x, n, limit float64) float64 {
    if v := math.Pow(x, n); v < limit {
        return v
    }
    return limit
}
```

### Switch statement
Go only runs the selected case, not all the cases that follow (the `break` statement is provided automatically)

Switch without a condition is the same as `switch true`. It can be Clean way to write long if-then-else chains:
```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```