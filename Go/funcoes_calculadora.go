package main

import (
	"fmt"
	"math"
)

var num1 float32
var num2 float32

func main() {

	option := 0

	for option != 3 {
		operation := 0
		menuPrincipal()
		fmt.Scan(&option)

		if option == 1 {
			menuCalcSimples()
			fmt.Scan(&operation)
			switch operation {
			case 1:
				informarDados()
				fmt.Printf("o Resultado da soma de %f e %f é : %2.f \n ", num1, num2, somar(num1, num2))
				break
			case 2:
				informarDados()
				fmt.Printf("o Resultado da subtração de %f e %f é : %2.f \n", num1, num2, subtrair(num1, num2))
				break
			case 3:
				informarDados()
				fmt.Printf("o Resultado da subtração de %f e %f é : %2.f \n", num1, num2, multiplicar(num1, num2))
				break
			case 4:
				informarDados()
				fmt.Printf("o Resultado da divisão de %f e %f é : %f \n", num1, num2, dividir(num1, num2))
				break
			case 5:
				fmt.Println("Escolheu sair bye!!!")
				break
			default:
				fmt.Println("opção invalida !!! Tente novamente")
				break
			}

		} else if option == 2 {
			menuCalcCientifica()
			fmt.Scan(&operation)
			switch operation {
			case 1:
				optPotencia()
				fmt.Printf("\nA potência de %f elevado a %f é : %2.f \n", num1, num2, calcularPotencia())
				break
			case 2:
				optRaizQuadrada()
				fmt.Printf("\nA raiz quadrada de %f é : %2.f\n", num1, calcularRaizQuadrada())
				break
			case 3:
				var valor int
				fmt.Println("Informe o valor:")
				fmt.Scan(&valor)
				fmt.Printf("O logaritmo de %d é %5.f\n", valor, math.Log(float64(valor)))
				break
			case 4:
				fmt.Println("Saindo")
			default:
				fmt.Println("Opção invalida, tente novamente.")
			}
		} else if option == 3 {
			fmt.Println("Saindo")
			break
		} else {
			fmt.Println("Opção inválida!")
		}
	}
}

func informarDados() {
	fmt.Print("Informe o 1º valor: ")
	fmt.Scan(&num1)
	fmt.Print("Informe o 2º valor: ")
	fmt.Scan(&num2)
}

func menuPrincipal() {
	fmt.Println("\n\n********************CALCULADORA********************")
	fmt.Println("1 - CALCULADORA SIMPLES ")
	fmt.Println("2 - CALCULADORA CIENTIFICA ")
	fmt.Println("3 - SAIR ")
	fmt.Println("Escolha a operação: ")
}

func menuCalcSimples() {
	fmt.Println("******************** CALCULADORA SIMPLES ********************")
	fmt.Println("1 - Somar ")
	fmt.Println("2 - Subtrair ")
	fmt.Println("3 - Multiplicar")
	fmt.Println("4 - Dividir ")
	fmt.Println("5 - Sair ")
}

func menuCalcCientifica() {
	fmt.Println("******************** CALCULADORA CIENTIFICA ********************")
	fmt.Println("1 - Potência ")
	fmt.Println("2 - Raiz quadrada ")
	fmt.Println("3 - Logaritmo")
	fmt.Println("4 - Sair ")
}

func somar(num1 float32, num2 float32) float32 {
	return num1 + num2
}

func subtrair(num1 float32, num2 float32) float32 {
	return num1 - num2
}

func multiplicar(num1 float32, num2 float32) float32 {
	return num1 * num2
}

func dividir(num1 float32, num2 float32) float32 {
	return num1 / num2
}

func calcularPotencia() float64 {
	return math.Pow(float64(num1), float64(num2))
}

func calcularRaizQuadrada() float32 {
	resultado := (num1 / 1.0 / 2.0)
	return resultado
}

func optPotencia() {
	fmt.Print("deseja calcular a potência de :")
	fmt.Scan(&num1)
	fmt.Print("\nelevado á :")
	fmt.Scan(&num2)
}

func optRaizQuadrada() {
	fmt.Print("Calcular a raiz Quadrada de :")
	fmt.Scanf("%f", &num1)
}
