package main
import "fmt"

type Usuario struct{
	Nombre     string
	Apellido   string
	Edad       int 
	Correo     string
	Contraseña string
}

func( u*Usuario) CambiarNombre(cambiarNombre, cambiarApellido string){
	u.Apellido = cambiarApellido
	u.Nombre = cambiarNombre
}

func(u*Usuario) CambiarEdad(nuevaEdad int ){
	u.Edad = nuevaEdad 
}

func(u*Usuario) CambiarCorreo(nuevoCorreo string){
	u.Correo = nuevoCorreo
}

func(u*Usuario) CambiarContraseña(nuevaContraseña string ){
	u.Contraseña = nuevaContraseña
}
func (u*Usuario) Mostrar(){
	fmt.Println("Nombre:", u.Nombre)
	fmt.Println("Apellido:", u.Apellido)
	fmt.Println("edad:", u.Edad)
	fmt.Println("correo:", u.Correo)
	fmt.Println("contraseña:", u.Contraseña)
}

func main() {
	// Creación de un usuario inicial
	usuario1 := Usuario{
		Nombre:     "jesus",
		Apellido:   "Pérez",
		Edad:       25,
		Correo:     "juan.perez@example.com",
		Contraseña: "clave123",
	}

	fmt.Println("Datos iniciales del usuario:")
	usuario1.Mostrar()

	// Modificando los datos usando las funciones
	usuario1.CambiarNombre("jesus", "castro")
	usuario1.CambiarEdad(24)
	usuario1.CambiarCorreo("anachurycastro2001@gmail.com")
	usuario1.CambiarContraseña("anachury456")

	fmt.Println("\nDatos después de las modificaciones:")
	usuario1.Mostrar()
}