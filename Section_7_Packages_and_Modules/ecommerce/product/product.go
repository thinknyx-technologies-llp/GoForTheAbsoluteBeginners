package product
import "fmt"

type Product struct{
	ID int
	Name string
	Price float64
}

var Products []Product

func AddProduct(id int, name string, price float64){
	Products = append(Products, Product{ID: id, Name: name, Price: price})
	fmt.Println("Product added successfully!")
}

func DisplayProducts(){
	fmt.Println("Product Catalogue: ")
	for _, p := range Products{
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}