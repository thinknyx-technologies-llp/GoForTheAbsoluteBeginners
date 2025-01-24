package main

import "fmt"

func demo2() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "Hello"
		ch1 <- "Thinknyx"
		ch1 <- "Technologies"
		close(ch1)
	}()
	for val := range ch1 {
		fmt.Println(val)
	}
}
