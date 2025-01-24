package main
import ("fmt"
    "os")
func demo2(){
    file, errs := os.Create("myfile.txt")
    _, errs = file.WriteString("Hello World!")
    data, errs := os.ReadFile("myfile.txt")
    fmt.Printf("%s", data)
    fmt.Println(errs)
}