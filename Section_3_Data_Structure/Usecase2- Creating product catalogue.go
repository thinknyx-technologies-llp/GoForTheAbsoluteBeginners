package main
import "fmt"

type Product struct {
	ID int
	Name string
	Price float64
}

func main (){
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 50000},
		{ID: 2, Name: "Smartphone", Price: 20000},
	}
	fmt.Println("Product Catalogue: ")
	for _, p:= range products{
		fmt.Printf("ID: %d, name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
	var id int
	var name string
	var price float64
	fmt.Println("\n Add a new product: ")
	fmt.Println("Enter ID: ")
	fmt.Scanln(&id)
	fmt.Println("Enter name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter price: ")
	fmt.Scanln(&price)
	products = append(products, Product{ID:id, Name:name, Price:price})
	fmt.Println("Product added successfully!")
	fmt.Println("############################################")
	fmt.Println("updated product catalogue:")
	for _, p:= range products{
		fmt.Printf("ID: %d, name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}