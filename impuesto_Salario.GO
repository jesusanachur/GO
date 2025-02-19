package main

import ("fmt")

// Función para calcular el impuesto sobre el salario
func calcularImpuesto(salario float64) float64 {
	var impuesto float64

	if salario > 150000 {
		impuesto = salario * 0.27 // 27% para salarios mayores a 150,000
	} else if salario > 50000 {
		impuesto = salario * 0.17 // 17% para salarios mayores a 50,000
	} else {
		impuesto = 0.0 // No hay impuesto para salarios menores o iguales a 50,000
	}

	return impuesto
}

func main() {
	var salario float64

	fmt.Print("Ingrese el salario del empleado: ")
	fmt.Scanln(&salario)

	impuesto := calcularImpuesto(salario)

	if impuesto > 0 {
		fmt.Printf("El impuesto a descontar del salario es: %.2f\n", impuesto)
	} else {
		fmt.Println("No se aplica impuesto para este salario.")
	}
}