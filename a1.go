package main

import (
	"fmt"
	"regexp"
)

func main() {
	txt := "The rain in Spain"
	pattern := "^The.*Spain$"

	matched, _ := regexp.MatchString(pattern, txt)

	if matched {
		fmt.Println("Yes we have a match")
	} else {
		fmt.Println("No match")
	}

}
