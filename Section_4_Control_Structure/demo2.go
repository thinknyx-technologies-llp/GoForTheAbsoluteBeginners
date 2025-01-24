package main
import "fmt"
func demo2(){
	var value interface{} = 42
	switch v:= value.(type){
	case int:
		fmt.Printf("Value is an integer: %d\n", v)
	case string:
		fmt.Printf("Value is a string: %s\n", v)
	default:
		fmt.Println("Unknown type!")
	}
}