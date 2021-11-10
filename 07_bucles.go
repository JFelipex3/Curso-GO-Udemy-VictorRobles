package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Bloque FOR
	for i := 1; i <= 3; i++ {
		if i%2 == 0 {
			fmt.Println("El Número " + strconv.Itoa(i) + " es Par")
		} else {
			fmt.Println("El Número " + strconv.Itoa(i) + " es Impar")
		}
	}

	fmt.Println("")

	// Recorrer Slice
	peliculas := []string{
		"Avengers End Game",
		"El Club de la Pelea",
		"Origen",
		"Gran Torino"}

	fmt.Println("Listado de Peliculas")
	fmt.Println("********************")
	for i := 0; i < len(peliculas); i++ {
		fmt.Println(strconv.Itoa(i+1) + ". " + peliculas[i])
	}

	fmt.Println("")

	// Simular For Each
	fmt.Println("Listado de Peliculas - Foreach")
	fmt.Println("********************")
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
}
