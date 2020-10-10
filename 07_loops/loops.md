# for loops
For is Go's "while", for example:
```go
sum := 1
for sum < 1000 {
	sum += sum
}
fmt.Println(sum)
```