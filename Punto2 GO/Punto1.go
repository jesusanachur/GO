package main

import ("fmt")

func main() {
    // se define la palabra 
    palabra := "hola mundo"

    // Contar todas las letras 
    Letras := len(palabra)
    
    // Imprimir todas las cantidades de letras
    fmt.Printf("La palabra  tiene", palabra, Letras)
    
    // Imprimir cada letra por separado
    fmt.Println("Las letras de la palabra son:")
    for _, letra := range palabra {
        fmt.Printf(" ", letra)
    }
}
