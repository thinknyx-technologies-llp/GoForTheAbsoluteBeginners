package main
import "fmt"

func type_switch(inter interface{}){
	switch value:= inter.(type){
	case int:
		fmt.Println("It is an int value")
	case string:
		fmt.Println("It is a string value")
	default:
		fmt.Printf("Unknown Data type, %t", value)
	}
}

func main(){
	type_switch("Thinknxy")
}