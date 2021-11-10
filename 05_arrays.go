package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Opción 1 declaranción de Arrays
	var peliculas [3]string
	peliculas[0] = "La verdad duele"
	peliculas[1] = "Ciudadano Ejemplar"
	peliculas[2] = "Gran Torino"

	fmt.Println(peliculas)

	// Opción 2 declaranción de Arrays
	peliculas2 := [3]string{"Spiderman", "Avengers: Era de Ultron", "Batman Inicia"}
	fmt.Println(peliculas2)

	// Array multidimensional
	var matrizPeliculas [3][2]string
	matrizPeliculas[0][0] = "La verdad duele"
	matrizPeliculas[0][1] = "Mientras duermes"

	matrizPeliculas[1][0] = "Ciudadano Ejemplar"
	matrizPeliculas[1][1] = "El Señor de los anillos"

	matrizPeliculas[2][0] = "Gran Torino"
	matrizPeliculas[2][1] = "Rápido y Furioso"

	fmt.Println(matrizPeliculas)
	fmt.Println(matrizPeliculas[0][1])

	// Slice es igual que un array, pero no tiene un úmero de indice determinado
	peliculasSlice := []string{
		"El Silencio de los Inocentes",
		"Matrix",
		"El Origen"}

	peliculasSlice = append(peliculasSlice, "Sin Límites")
	peliculasSlice = append(peliculasSlice, "Camp Rock")

	fmt.Println(peliculasSlice)
	fmt.Println("Largo Slice: " + strconv.Itoa(len(peliculasSlice)))
	fmt.Println("Obtener valores del 0 al 3 (no inclusivo)")
	fmt.Println(peliculasSlice[0:3])
}
