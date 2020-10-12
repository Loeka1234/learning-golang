# Packages

- Every Go program is made up of packages

- Program starts running in package `main`

- `math/rand` package:
  
  - `rand.Intn(10)` always returns the same number
  
  - Solution: 
    
    - Seed the number generator:
    
    - ```go
      rand.Seed(time.Now().UnixNano())
      ```