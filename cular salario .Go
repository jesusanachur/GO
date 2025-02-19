package main
import ("fmt")

//  para calcular el salario basado en minutos trabajados y categoría
func calcularSalario(minutosTrabajados int, categoria string) float64 {
	var salarioBase float64

	// Convertir minutos a horas
	horasTrabajadas := float64(minutosTrabajados) / 60.0

	//  el salario base según la categoría
	switch categoria {
	case "C":
		salarioBase = 1000.0
	case "B":
		salarioBase = 1500.0
	case "A":
		salarioBase = 3000.0
	default:
		fmt.Println("Categoría no válida")
		return 0.0
	}

	// Calcular el salario mensual
	salarioMensual := salarioBase * horasTrabajadas

	//  el porcentaje adicional según la categoría
	switch categoria {
	case "B":
		salarioMensual += salarioMensual * 0.20 // 20% adicional
	case "A":
		salarioMensual += salarioMensual * 0.50 // 50% adicional
	}

	return salarioMensual
}

func main() {
	var categoria string
	var minutosTrabajados int

	fmt.Print("Ingrese la categoría del empleado (C, B, A): ")
	fmt.Scanln(&categoria)
	fmt.Print("Ingrese la cantidad de minutos trabajados en el mes: ")
	fmt.Scanln(&minutosTrabajados)

	salario := calcularSalario(minutosTrabajados, categoria)

	if salario > 0 {
		fmt.Printf("El salario mensual del empleado es: %.2f\n", salario)
	}
}