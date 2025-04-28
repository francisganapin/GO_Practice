package main
import "fmt"

func main() {
    nums := []int{10, 20, 30} // Variable name fixed
    for index, value := range nums {
        fmt.Println("Index:", index, "Value:", value) // fmt is now used
    }
}