package main
import "fmt"
func demo3(){
	empName:=[3]string{"Aryan", "Dheeraj", "Kanishk"}
	for idx, val:= range empName{
		fmt.Printf("%v\t%v\n", idx, val)
	}
}