package main

import (
	"fmt"
	"regexp"
)

func main() {
	txt := "The rain in Spain"
	pattern := "^The. *Spain$"

	re := regexp.MustCompile(pattern)

	matched := re.FindAllString(txt, -1)

	fmt.Println(matched)
}
