package main
import("fmt")

type producto struct{
	id int
	nombre string
	prince float64
	description string
	categoria string
}
var products [] producto
func (p * producto) save () {
	products  = append(products, * p)
}
func (p*producto ) GeTAll(){
	for _, producto:= range products {
		fmt.Printf("id: %d, nombre: %s, Price: %.2f, description: %s, Categoria: %s\n", producto.id, producto.nombre, producto.prince, producto.description, producto.categoria)
		
	}
}
func getById(id int) *producto {
	for _, product := range products {
		if product.id == id {
			return &product
		}
	}
	return nil
}


func main() {
	// Crear algunos productos
	product1 := producto{id: 1, nombre: "Laptop", prince: 1200.50, description: "High-end gaming laptop", categoria: "Electronics"}
	product2 := producto{id: 2, nombre: "Smartphone", prince: 800.00, description: "Latest model smartphone", categoria: "Electronics"}package main
	import "fmt"
	
	type producto struct {
		id          int
		nombre      string
		prince      float64
		description string
		categoria   string
	}
	
	var products []producto
	
	func (p *producto) save() {
		products = append(products, *p)
	}
	
	func (p *producto) GeTAll() {
		for _, producto := range products {
			fmt.Printf("id: %d, nombre: %s, Price: %.2f, description: %s, Categoria: %s\n", producto.id, producto.nombre, producto.prince, producto.description, producto.categoria)
		}
	}
	
	func getById(id int) *producto {
		for _, product := range products {
			if product.id == id {
				return &product
			}
		}
		return nil
	}
	
	func main() {
		// Crear algunos productos
		product1 := producto{id: 1, nombre: "Laptop", prince: 1200.50, description: "High-end gaming laptop", categoria: "Electronics"}
		product2 := producto{id: 2, nombre: "Smartphone", prince: 800.00, description: "Latest model smartphone", categoria: "Electronics"}
	
		// Guardar los productos en el slice products
		product1.save()
		product2.save()
	
		// Imprimir todos los productos
		product1.GeTAll()
	
		// Obtener un producto por ID
		foundProduct := getById(1)
		if foundProduct != nil {
			fmt.Printf("Product found: %s\n", foundProduct.nombre)
		} else {
			fmt.Println("Product not found")
		}
	}
	// Guardar los productos en el slice Products
	products.save()()
	products.save()()

	// Imprimir todos los productos
	products.GetAll()

	// Obtener un producto por ID
	foundProduct := getById(1)
	if foundProduct != nil {
		fmt.Printf("Product found: %s\n", foundProduct.nombre)
	} else {
		fmt.Println("Product not found")
	}
}