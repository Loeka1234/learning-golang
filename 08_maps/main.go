package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["bob"] = "bob@bobie.com"
	emails["bert"] = "bert@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])

	// Delete from map
	delete(emails, "bob")

	fmt.Println(emails)

	// Declare map and add key values
	scores := map[string]int{"Loeka": 8, "Bob": 3}

	fmt.Println(scores)
}
