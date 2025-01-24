package main
import "fmt"
func demo1(){
	empAge:= 42
	if empAge > 40{
		fmt.Println("Employee is eligible for PF.")
		if empAge > 45{
			fmt.Println("Employee is also eligible for retirement.")
		}
	}
}