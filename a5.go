package main

import "fmt"

func main() {
	car := []string{"yamaha", "Toyato", "Camaro"}

	for _, x := range car {
		fmt.Println(x)
	}
}
