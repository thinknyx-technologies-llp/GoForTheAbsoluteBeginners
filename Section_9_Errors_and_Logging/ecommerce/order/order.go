package order

import (
	"fmt"
	"log"
	"os"
)

type Order struct {
	ID       int
	Product  string
	Quantity int
}

var Orders []Order

func PlaceOrder(id int, product string, quantity int) {
	Orders = append(Orders, Order{ID: id, Product: product, Quantity: quantity})
	fmt.Println("Order placed successfully!")
	logOrderToFile(Order{ID: id, Product: product, Quantity: quantity})
}

func DisplayOrders() {
	fmt.Println("Order History:")
	for _, o := range Orders {
		fmt.Printf("Order ID: %d, Product: %s, Quantity: %d\n", o.ID, o.Product, o.Quantity)

	}
}

func logOrderToFile(order Order){
	file, err := os.OpenFile("orders.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()
	logger := log.New(file, "ORDER_LOG:", log.LstdFlags)
	logger.Printf("Order ID: %d, Product: %s, Quantity: %d\n", order.ID, order.Product, order.Quantity)
}