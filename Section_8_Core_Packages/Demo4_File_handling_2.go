package main
import (
	"fmt"
	"os")
func demo3(){
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) // os.O_WRONLY is write-read only mode
	if err != nil{
		fmt.Println(err)
		return
	}
	_, err = file.WriteString("This is the appended line.")
	if err != nil{
		fmt.Println(err)
	}
}
