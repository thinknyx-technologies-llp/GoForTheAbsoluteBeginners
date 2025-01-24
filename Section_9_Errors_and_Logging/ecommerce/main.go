package main
import (
	"ecommerce/product"
	"ecommerce/user"
	"ecommerce/order"
	"fmt"
)

func main(){
	user1:= user.Register("admin", "password")
	if user.Login(user1, "admin", "Password"){
	fmt.Println("Login successfylly")
	} else{
		fmt.Println("Invalid Credentials!")
	}
	product.AddProduct(1, "Laptop", 50000)
	product.AddProduct(2, "Smartphone", 20000)
	product.DisplayProducts()
	order.PlaceOrder(1, "Laptop", 1)
	order.DisplayOrders()
} 