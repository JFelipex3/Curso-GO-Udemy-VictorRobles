package main

import (
	"fmt"
	"time"
)

func main() {

	var suma int = 8 + 9
	fmt.Println(suma)

	var resta int = 6 - 4
	fmt.Println(resta)

	var nombre string = "Jhonnatan"
	fmt.Println("Hola Mundo!!! " + nombre)

	var apellidos string = "Machuca"
	fmt.Println("Hola Mundo!!! " + nombre + " " + apellidos)

	pais := "Chile" // Define automaticamente el tipo de datos a la variable
	fmt.Println(pais)

	var existe bool = true
	fmt.Println(existe)

	var decimal float32 = 12.24
	fmt.Println(decimal)

	const year int = 2021 // Corresponde a una constante que no se puede modificar
	fmt.Println(year)

	time.Sleep(time.Second * 5)               // Retarda la ejecución en 5 segundos
	fmt.Println("Luego de los 5 Segundos...") // Librería fmt permite escribir texto
}
