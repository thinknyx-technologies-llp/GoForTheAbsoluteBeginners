package main
import "fmt"
func demo3(){
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{6,7,8,9}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}