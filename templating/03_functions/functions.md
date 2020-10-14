# Functions

## `and` functions
```go
{{if and .User .User.Admin}}
    You are an admin user!
{{else}}
    Access denied!
{{end}}
```
Equivalent to `if a then b else a`

## Comparison functions
- `eq` - Returns the boolean truth of `arg1 == arg2`
- `ne` - Returns the boolean truth of `arg1 != arg2`
- `lt` - Returns the boolean truth of `arg1 < arg2`
- `le` - Returns the boolean truth of `arg1 <= arg2`
- `gt` - Returns the boolean truth of `arg1 > arg2`
- `ge` - Returns the boolean truth of `arg1 >= arg2`