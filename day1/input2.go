package main

import "fmt"

func main() {
	var age int
	var city string

	fmt.Println("Enter your city: ")
	fmt.Scan(&city)
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Age", age, "City", city)
}
