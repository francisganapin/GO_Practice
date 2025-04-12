package main

import (
	"fmt"
)

func famlyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "refsnes")
}

func main() {
	famlyName("lIAM", 3)
	famlyName("jenny", 14)
	famlyName("anja", 30)
}
