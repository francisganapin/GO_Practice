package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	exists := false
	for _, fruit := range fruits {
		if fruit == "banana" {
			exists = true
			break
		}
	}
	if exists {
		fmt.Println("banana exist")
	} else {
		fmt.Println("No banana not found")
	}
}
