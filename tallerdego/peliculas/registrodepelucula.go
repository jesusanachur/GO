
package main
import ("fmt")

type Pelicula struct {
    titulo      string
    director    string
    estreno  int
    Rating    float32
}

func actualizarRating(p *Pelicula, nuevoRating float32) {
	p.Rating =  nuevoRating 

}


func mostrarPelicula(p Pelicula) string {
	fmt.Println("La pelicula es: ", p.titulo,p.director, p.estreno, p.Rating)

		
}
func main() {
    pelicula := Pelicula{
        titulo:   "El rambo ",
        director: "Francis Ford Coppola",
        estreno:     1972,
        Rating:   9.2,
    }
    
    actualizarRating(&pelicula, 9.5)
    mostrarPelicula(pelicula)
}