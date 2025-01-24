package main
import "fmt"

func empDetails(name string, age int)(result string, age_result int){
	result = name
	age_result = age
	return
}

func demo5(){
	fmt.Println(empDetails("kanishk", 24))
}