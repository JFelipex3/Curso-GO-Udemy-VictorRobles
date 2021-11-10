package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Lector: ")

	nuevo_texto := os.Args[1] + "\n"

	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND, 0777)
	showError(err)
	escribir, err := fichero.WriteString(nuevo_texto)
	showError(err)
	fmt.Println(escribir)

	fichero.Close()

	contenido, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	texto := string(contenido) // Usado para convertir de byte a texto
	fmt.Println(texto)
}

func showError(e error) {
	if e != nil {
		panic(e) // Permite cortar la ejecución y mostrar error
	}
}
