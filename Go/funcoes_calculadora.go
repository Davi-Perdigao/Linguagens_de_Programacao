package main

import (
	"fmt"
)

func main() {
	var valor1 float64
	var valor2 float64
	var op string

	fmt.Println("Informe os valores:")
	fmt.Scan(&valor1)
	fmt.Scan(&valor2)

	fmt.Println("Informe o tipo de operação:")
	fmt.Scan(&op)

	//fmt.Println(valor1, valor2, op)

	switch op {
	case "+":
		fmt.Println(somar(valor1, valor2))
		break
	case "-":
		fmt.Println(subtrair(valor1, valor2))
		break
	case "*":
		fmt.Println(multiplicar(valor1, valor2))
		break
	case "/":
		fmt.Println(dividir(valor1, valor2))
		break
	default:
		fmt.Println("Informe um valor válido")
		break
	}
}

func somar(valor1 float64, valor2 float64) float64 {
	return valor1 + valor2
}

func subtrair(valor1 float64, valor2 float64) float64 {
	return valor1 - valor2
}

func multiplicar(valor1 float64, valor2 float64) float64 {
	return valor1 * valor2
}

func dividir(valor1 float64, valor2 float64) float64 {
	return valor1 / valor2
}
