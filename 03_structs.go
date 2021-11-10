package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {

	var gorraNegra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}

	// var gorraNegra = Gorra{"Nike", "Negro", 25.50, false} // Solo cuando se sabe el orden de las propiedades

	fmt.Println(gorraNegra)
	fmt.Println(gorraNegra.marca)
}
