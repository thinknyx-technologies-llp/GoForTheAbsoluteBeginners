package main

import (
	"fmt"
	"time"
)

func demo1() {
	go func(str string){
		fmt.Println(str)
		time.Sleep(500 * time.Millisecond)
	}("Hello Thinknyx")
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Main function")
}
