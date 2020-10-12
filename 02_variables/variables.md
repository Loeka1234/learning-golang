# Variables
### Variable Declaration

```go
var foo int
var foo int = 42
var foo, bar, int = 131
foo := 42
foo, bar := true, false
```

Can't redeclare variables, but can shadow them
All variables must be used

### Basic types

- Types
  
  - bool
  
  - string
  
  - int, int8, int16, int32, int64, uint, uint8, uint 16, uint 32, uint64, uintptr 
  
  - byte (alias for uint8)
  
  - rune (alias for int32, represents a Unicode code point)
  
  - float32, float64
  
  - complex64, complex128

- Info
  
  - `int`, `uint` & `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems

### Zero values

Variables declared without an explicit initial value are given their zero value

- 0 for numeric types

- false for boolean type

- "" (empty string) for strings

### Type conversions

T(v) converts v to type T

```go
x := 42
y := float64(y)
z := uint(y)
```

Use `strconv` package for strings

### Constants

- Declared using `const` keyword: `const hello = "world"`

- Cannot be declared using `:=` syntax

- Can be character, string, boolean, or numeric