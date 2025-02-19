package main

import ("fmt")



func main() {
    // Definir la estructura del cliente
    cliente := Cliente{
        Nombre:     "Juan Pérez",
        Edad:       25,
        Empleado:   true,
        Antiguedad: 2,
        Sueldo:     120000,
    }

    // Evaluar la elegibilidad del cliente
    elegible, tasaInteres := esElegible(cliente)
    
    if Edad >= 22 && Empleado  && Antiguedad >= 1  {
        if Sueldo >100000 {
            fmt.Printf("El cliente  es elegible para un préstamo sin intereses.\n", )
        } else {
            fmt.Printf("El cliente %s es elegible para un préstamo con una tasa de interés del .\n")
        }
    } else {
        fmt.Printf("El cliente no es elegible para un préstamo.\n",)
    }
}
