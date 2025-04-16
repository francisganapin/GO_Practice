package main

import (
	"fmt"
	"regexp"
)

func main() {
	txt := "the rain in Spain"
	pattern := "\\s+"

	re := regexp.MustCompile(pattern)
	words := re.Split(txt, -1)

	fmt.Println(words)
}
