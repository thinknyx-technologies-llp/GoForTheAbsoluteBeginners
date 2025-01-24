package main

import (
	"fmt"
	"time"
)

func square_number1 (number int){
	time.Sleep(1 * time.Second)
	fmt.Println(number * number)
}

func demo3(){
	initial_time := time.Now()
	for i:=1;i<=1000;i++{
		go square_number1(i)
	}
	end_time:= time.Since(initial_time)
	time.Sleep(2 * time.Second)
	fmt.Println(end_time)
}