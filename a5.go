package main
import "fmt"


func main(){
	thisDict := map[string]interface{}{
		"brand":"ford",
		"model":"mustang",
		"year":1964,
	}
	fmt.Println(len(thisDict))
}