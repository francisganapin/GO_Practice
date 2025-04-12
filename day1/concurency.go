package main

import(
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello,World")
}

func main(){
	go sayHello()
	time.Sleep(time.Second)
	fmt.Println("Main function finished")
}