package main

import (
	"awesomeProject/Calculadora"
	"awesomeProject/Gerador_senhas"
	"awesomeProject/Jogo_da_forca"
	"fmt"
)

func main() {

	for {
		var escolha int
		quebralinha()
		fmt.Println("Escolha uma funçao\nCalduladora =>1\nGerador de senhas =>2\nJogo da forca =>3\nSair =>4")
		quebralinha()
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			Calculadora.Calculadora()
		case 2:
			Gerador_senhas.Gerador()
		case 3:
			Jogo_da_forca.Jogo_da_forca()
		case 4:
			return
		default:
			fmt.Println("Escola invalida!")
		}
	}

}
