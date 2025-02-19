package main
import ("fmt")

//  para calcular el promedio de un slice de enteros
func calcularPromedio(notas []int) float64 {
	var suma int
	var contador int

	for _, nota := range notas {
		if nota < 0 {
			fmt.Println("Error: No se permiten notas negativas.")
			return 0.0
		}
		suma += nota
		contador++
	}

	if contador == 0 {
		return 0.0 // Evitar división por cero
	}

	return float64(suma) / float64(contador)
}

func main() {
	var notas []int
	var nota int

	fmt.Println("Ingrese las notas (ingrese un número negativo para terminar):")

	for {
		fmt.Print("Nota: ")
		fmt.Scanln(&nota)

		if nota < 0 {
			break // Salir del bucle si se ingresa una nota negativa
		}

		notas = append(notas, nota) // Agregar la nota al slice
	}

	promedio := calcularPromedio(notas)

	if promedio > 0 {
		fmt.Printf("El promedio de las notas es: %.2f\n", promedio)
	} else {
		fmt.Println("No se ingresaron notas válidas.")
	}
}