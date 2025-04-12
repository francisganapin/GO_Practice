package main

import ("fmt")

func main(){
	fruits := map[string]int{
		"apple":500,
		"bananas":300,
		"orange":200,
	}

	numb := fruits["apple"]

	fmt.Printf("There are %d apple .\n",numb)
}