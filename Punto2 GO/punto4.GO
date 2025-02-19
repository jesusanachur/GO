package main

import ("fmt")

// la funcion del  para calcular el precio con descuento
func Descuento(precio , descuento )  {
    return precio * (1 - descuento/100)
}

func main() {
    //se  Definir el precio y el porcentaje de descuento
    precio := 100.0 // Precio del producto
    descuento := 20.0 // Porcentaje de descuento

    // Calcular el precio con descuento
    precioConDescuento := calcularDescuento(precio, descuento)

    // Imprimir el resultado
    fmt.Printf("El precio original es", precio)
    fmt.Printf("El descuento es ", descuento)
    fmt.Printf("El precio con descuento es", precioConDescuento)
}
