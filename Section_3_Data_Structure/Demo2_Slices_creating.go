package main
import "fmt"
func demo2(){
	nums := make([] int, 3, 5)
	fmt.Println("slice: ", nums)
	fmt.Println("length: ", len(nums))
	fmt.Println("capacity: ", cap(nums))
}