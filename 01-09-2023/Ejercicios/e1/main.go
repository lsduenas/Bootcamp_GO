package main

import "fmt"


type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category string
}

var products = []Product {
	{
		ID: 123,
		Name: "Arroz",
		Price: 34.9,
		Description: "Arroz Roa",
		Category: "cereal",
	},
	{
		ID: 456,
		Name: "Lentejas",
		Price: 23.9,
		Description: "La soberana",
		Category: "granos",
	},
}

func (p Product) Save(){
	fmt.Println("Product saved")
	fmt.Println(append(products, p))
}

func (p Product) GetAll(){
	fmt.Println(
		"Listado de productos",
		products,
	)
}

func getById(ID int)(p Product){
	for _, key:= range products {
		if key.ID == ID {
			p = key
		}
	}
	return
}

func main(){
	p1:= Product {
		ID: 789,
		Name: "Pasta",
		Price: 21.4,
		Description: "La muñeca",
		Category: "cereal",
	}
	// añadir nuevo producto al slice de products
	p1.Save()

	// Mostrar productos
	//products = append(products, p1)
	p1.GetAll()

	// Get product by ID
	p2:= getById(456)
	fmt.Println("Product By ID", p2)

}