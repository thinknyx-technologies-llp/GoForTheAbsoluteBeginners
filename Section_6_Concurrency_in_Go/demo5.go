package main
import _"fmt"
func demo5(){
	channel := make(chan int, 10)
	channel <- 100
	close(channel)
	close(channel)
}