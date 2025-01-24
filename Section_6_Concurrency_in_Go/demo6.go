package main

import (
	"fmt"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func processOrder(orderID int, product Product, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintf("Order %d for %s processed.", orderID, product.Name)
}

func main() {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 50000},
		{ID: 2, Name: "Smartphone", Price: 20000},
	}
	var wg sync.WaitGroup
	orderChannel := make(chan string)
	fmt.Println("Placing Orders...")
	for i, product := range products {
		wg.Add(1)
		go processOrder(i+1, product, &wg, orderChannel)
	}
	go func() {
		wg.Wait()
		close(orderChannel)
	}()
	for status := range orderChannel {
		fmt.Println(status)
	}
}
