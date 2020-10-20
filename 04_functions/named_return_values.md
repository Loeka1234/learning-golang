# Named return values

- Known as "naked" return

- Should be used only in short functions., they can harm readability in longer functions

```go
func numbers() (x, y int) {
    x = 7
    y = 1
    return
}
```