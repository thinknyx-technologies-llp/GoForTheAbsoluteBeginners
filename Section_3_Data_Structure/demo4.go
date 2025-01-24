package main
import "fmt"
func demo4(){
	var empDetails = make(map[string]string)
    empDetails["name"] = "kanishk"
    empDetails["age"] = "23"
    empDetails["designation"] = "project engineer"
    delete(empDetails, "age")
    fmt.Println(empDetails)
}