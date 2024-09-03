package Calculadora

import (
	"fmt"
	"strings"
)

func quebralinha() {
	fmt.Println("=========================================================")
}

func Calculadora() {
	for {
		var num1, num2, res float32
		var opera int
		var sair string

		quebralinha()
		fmt.Println("====================== Calculadora ======================")
		quebralinha()

		fmt.Print("Qual é o primeiro número?\n")
		fmt.Scan(&num1)
		quebralinha()

		fmt.Print("Qual é o segundo número?\n")
		fmt.Scan(&num2)
		quebralinha()

		fmt.Print("Escolha uma opção:\n")
		fmt.Println("Soma => 1")
		fmt.Println("Subtração => 2")
		fmt.Println("Divisão => 3")
		fmt.Println("Multiplicação => 4")
		fmt.Println("Sair => 5")
		fmt.Scan(&opera)

		switch opera {
		case 1:
			res = num1 + num2
			fmt.Printf("A soma entre %f e %f é %f\n", num1, num2, res)
		case 2:
			res = num1 - num2
			fmt.Printf("A subtração entre %f e %f é %f\n", num1, num2, res)
		case 3:
			if num2 == 0 {
				fmt.Println("Não é possível dividir por zero!")
			} else {
				res = num1 / num2
				fmt.Printf("A divisão entre %f e %f é %f\n", num1, num2, res)
			}
		case 4:
			res = num1 * num2
			fmt.Printf("A multiplicação entre %f e %f é %f\n", num1, num2, res)
		case 5:
			fmt.Print("Voltando ao menu inical!")
			return
		default:
			fmt.Println("Erro! Opção inválida!")
		}

		fmt.Print("Deseja tentar novamente? ('s' ou 'n')\n")
		fmt.Scan(&sair)
		if strings.ToLower(sair) == "n" {
			break
		}
	}
}
