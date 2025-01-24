package main
import "fmt"

type empDetails struct{
	empName string
	empSalary int
	empPosition string
	empAge int
}

func demo5(){
	var emp1 empDetails
	var emp2 empDetails
	emp1.empName = "kanishk"
	emp1.empSalary = 6000
	emp1.empPosition = "Teacher"
	emp1.empAge = 24
	emp2.empName = "aryan"
	emp2.empSalary = 5000
	emp2.empPosition = "Intern"
	emp2.empAge = 22
	fmt.Println("name: ", emp1.empName)
	fmt.Println("salary: ", emp1.empSalary)
	fmt.Println("posotion: ", emp1.empPosition)
	fmt.Println("age: ", emp1.empAge)
	fmt.Println("name: ", emp2.empName)
	fmt.Println("salary: ", emp2.empSalary)
	fmt.Println("position: ", emp2.empPosition)
	fmt.Println("age: ", emp2.empAge)

}