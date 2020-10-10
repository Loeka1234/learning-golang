package main

import "fmt"

func main() {
	IDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Loop through IDs
	for i, id := range IDs {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without using index
	for _, id := range IDs {
		fmt.Println(id)
	}

	// All ids together
	sum := 0
	for _, id := range IDs {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	// Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "bert": "bert@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: ", key)
	}
}
