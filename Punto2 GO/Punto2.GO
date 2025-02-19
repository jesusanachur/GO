package main

import ("fmt")

func main() {
    // Número del mes
    numero:= 2
    meses := map[int]string{
        1: "Enero",
        2: "Febrero",
        3: "Marzo",
        4: "Abril",
        5: "Mayo",
        6: "Junio",
        7: "Julio",
        8: "Agosto",
        9: "Septiembre",
        10: "Octubre",
        11: "Noviembre",
        12: "Diciembre",
    }
    
    
    // Imprimir el nombre del mes
    fmt.Printf("El número  corresponde al mes ", numero, meses)
}

meses, ok := meses[numero]
if ok {
    fmt.Printf("El número  corresponde al mes ", numero, meses)
} else {
    fmt.Printf("El número no corresponde a ningún mes", numero)
}

