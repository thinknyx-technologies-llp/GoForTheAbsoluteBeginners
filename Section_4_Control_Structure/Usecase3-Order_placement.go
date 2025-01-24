package main
import "fmt"

type Product struct{
	ID int
	Name string
	Price float64
}

func placeOrder(productID int, products []Product){
	for _,p:= range products{
		if p.ID == productID{
			fmt.Printf("Order placed for a %s of %.2f\n", p.Name, p.Price)
			return
		}
	}
	fmt.Println("Invalid Product ID!")
}

func main(){
	products := []Product{
		{ID: 1, Name: "laptop", Price: 50000},
		{ID: 2, Name: "Smartphone", Price: 20000},
	}
	fmt.Println("Available Products: ")
	for _, p:= range products{
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
	var productID int
	fmt.Print("\nEnter product ID to place order: ")
	fmt.Scanln(&productID)
	placeOrder(productID, products)
}