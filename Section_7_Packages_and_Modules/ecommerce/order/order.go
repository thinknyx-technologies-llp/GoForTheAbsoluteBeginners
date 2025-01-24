package order
import "fmt"

type Order struct{
	ID int
	Product string
	Quantity int
}

var Orders []Order

func PlaceOrder(id int, product string, quantity int){
	Orders = append(Orders, Order{ID: id, Product: product, Quantity: quantity})
	fmt.Println("Order placed successfully!")
}

func DisplayOrders(){
	fmt.Println("Order History:")
	for _, o := range Orders{
		fmt.Printf("Order ID: %d, Product: %s, Quantity: %d\n", o.ID, o.Product, o.Quantity)
		
	}
}