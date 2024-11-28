package main

import (
	"fmt"
)

func main() {

	// Get the user's input Don't change this!
	var operation string
	var name string
	var value int
	fmt.Scanln(&operation)
	fmt.Scanln(&name)
	fmt.Scanln(&value)
	// Map to store the books in the library Don't change this!
	library := map[string]int{
		"Golang": 134,
		"Python": 200,
		"Rust":   300,
		"Lua":    260,
		"Java":   500,
	}
	// Check for operation Type (Add or Delete)
	if operation == "Add" {
		// Add the book with value
		library[name] = value
		fmt.Println("Book added successfully!")
		fmt.Println(library)
	} else {
		// Delete the book
		if operation == "Remove" {
			delete(library, name)
			fmt.Println("Book deleted successfully!")
			fmt.Println(library)
		}
	}
}
