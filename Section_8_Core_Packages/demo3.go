package main
import (
	"fmt"
	"os")
func demo3(){
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil{
		fmt.Println(err)
		return
	}
	_, err = file.WriteString("This is the appended line.")
	if err != nil{
		fmt.Println(err)
	}
}