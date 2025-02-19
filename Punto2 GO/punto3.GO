package main

import ("fmt")
func main() {
    // se define  la lista de  los estudiantes
    estudiantes := []string{"Benjamin", "Breda","Cristian", "Diana", "Eduardo", "Fernanda", "Gabriela", "Hugo",  "Javier", "Karla", "Luis", "Mariana", "Natalia", "Oscar", "Pamela", "Quetzalli", "Ricardo", "Sofia", "Tania", "Ulises", "Viviana", "Wendy", "Ximena", "Yahir", "Zaira"}

    // se  Imprimir la lista de estudiantes
    fmt.Println("Lista de estudiantes:")
    for i , estudiante := range estudiantes {
        fmt.Println(estudiante)
    }
	// se INGRESA UN NUEVO ESTUDIANTE
	estudianteNuevo := "Zacarias"
	fmt.Println("lista actulizada :")
	for i , estudiante := range estudiantes {
		fmt.Println(estudiante)
	}

}
