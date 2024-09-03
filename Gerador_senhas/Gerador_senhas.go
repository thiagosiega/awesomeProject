package Gerador_senhas

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func quebralinha() {
	fmt.Println("=========================================================")
}

func gerar(tamanho int, cactere bool, text bool) string {
	// Gerar alfabeto em minúsculas
	var alfabetoMinusculas []rune
	for c := 'a'; c <= 'z'; c++ {
		alfabetoMinusculas = append(alfabetoMinusculas, c)
	}

	// Gerar alfabeto em maiúsculas
	var alfabetoMaiusculas []rune
	for c := 'A'; c <= 'Z'; c++ {
		alfabetoMaiusculas = append(alfabetoMaiusculas, c)
	}

	// Gerar caracteres especiais
	var caracteresEspeciais []rune
	if cactere {
		caracteresEspeciais = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+'}
	}

	// Combinar todos os caracteres permitidos
	var caracteres []rune
	caracteres = append(caracteres, alfabetoMinusculas...)
	caracteres = append(caracteres, alfabetoMaiusculas...)
	if cactere {
		caracteres = append(caracteres, caracteresEspeciais...)
	}
	if text {
		// Adicionar números ao conjunto de caracteres se "text" for verdadeiro
		for c := '0'; c <= '9'; c++ {
			caracteres = append(caracteres, c)
		}
	}

	// Gerar a senha
	rand.Seed(time.Now().UnixNano())
	senha := make([]rune, tamanho)
	for i := range senha {
		senha[i] = caracteres[rand.Intn(len(caracteres))]
	}

	return string(senha)
}

func confirma() bool {
	var escolha string
	fmt.Println("Escolha uma opção:\nSim => s\nNão => n")
	fmt.Scan(&escolha)
	return strings.ToLower(escolha) == "s"
}

func Gerador() {
	for {
		var tamanho int
		var cactere, text bool
		var sair string

		quebralinha()
		fmt.Println("A senha deve ter quantos caracteres?")
		fmt.Scan(&tamanho)

		quebralinha()
		fmt.Println("A senha deve ter caracteres especiais? (Sim ou Não)")
		cactere = confirma()

		quebralinha()
		fmt.Println("A senha deve ter números? (Sim ou Não)")
		text = confirma()

		quebralinha()
		senha := gerar(tamanho, cactere, text)
		fmt.Println("Senha gerada:", senha)

		quebralinha()
		fmt.Println("Deseja sair?\nEscolha uma opção:\nSim => s\nNão => n")
		fmt.Scan(&sair)
		if strings.ToLower(sair) == "s" {
			fmt.Println("Voltand ao menu...")
			return
		}
	}

}
