package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("***** Mi Programa con Go *****")
	/*
		var edad int

		fmt.Println("Hola " + os.Args[1] + " bienvenido al programa en GO")
		edad, _ = strconv.Atoi(os.Args[2]) // retorno 2 valores, el valor convertido y un error que pueda ocurrir, se puede utilizar _ para no indicar una variable

		if edad >= 18 && edad <= 99 {
			fmt.Println("Eres mayor de edad")
		} else if edad > 99 {
			fmt.Println("Eres demasiado mayor")
		} else {
			fmt.Println("Eres menor de edad")
		}
	*/

	numero, _ := strconv.Atoi(os.Args[1])

	if numero%2 == 0 {
		fmt.Println("El Número es Par")
	} else {
		fmt.Println("El Número es Impar")
	}
}
