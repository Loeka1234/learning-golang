package main

import "fmt"

type Vertex struct {
	Let, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)

	// Inserting
	m2["Google"] = Vertex{1,2}
	// Retrieving
	_ = m2["Google"]
	// Test if key is present
	_, ok := m2["Google"]
	fmt.Println(ok)
	// Delete
	delete(m2, "Google")
}
