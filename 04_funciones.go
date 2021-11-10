package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var numero1, numero2 float32 = 10, 6

	fmt.Println("Calculadora #1")
	calculadora(numero1, numero2)
	fmt.Print("\n")

	numero1, numero2 = 44, 7
	fmt.Println("Calculadora #2")
	calculadora(numero1, numero2)
	fmt.Print("\n")

	fmt.Println("Pedido #1")
	fmt.Println(gorras(45, "$"))
	fmt.Print("\n")

	fmt.Println("Pedido #2")
	fmt.Println(gorras(24, "$"))
	fmt.Print("\n")

	fmt.Println("Pantalon #1")
	pantalon("Rojo", "largo", "sin bolsillos", "Nike")
	fmt.Print("\n")
}

func operacion(n1 float32, n2 float32, operador string) float32 {

	var resultado float32

	if operador == "+" {
		resultado = n1 + n2
	} else if operador == "-" {
		resultado = n1 - n2
	} else if operador == "*" {
		resultado = n1 * n2
	} else if operador == "/" {
		resultado = n1 / n2
	}

	return resultado
}

func calculadora(n1 float32, n2 float32) {
	fmt.Print("La suma es: ")
	fmt.Println(operacion(n1, n2, "+"))

	fmt.Print("La resta es: ")
	fmt.Println(operacion(n1, n2, "-"))

	fmt.Print("La multiplicación es: ")
	fmt.Println(operacion(n1, n2, "*"))

	fmt.Print("La división es: ")
	fmt.Println(operacion(n1, n2, "/"))
}

func gorras(pedido float32, moneda string) (string, float32, string) {

	precio := func() float32 {
		return pedido * 7
	}

	return "El precio de gorras pedidas es: ", precio(), moneda
}

func pantalon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}
