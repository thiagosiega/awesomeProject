package Jogo_da_forca

import (
	"fmt"
	"math/rand"
	"time"
)

func Jogo_da_forca() {
	var escola_palavra string
	lista_palavras := []string{"sol", "lua", "estrela", "planeta", "galáxia", "universo", "cometa", "asteroide", "satélite"}
	// inicia o gerador
	rand.Seed(time.Now().UnixNano())
	numero := rand.Intn(len(lista_palavras))
	palavra := lista_palavras[numero]
	progresso := make([]rune, len(palavra))
	for i := range progresso {
		progresso[i] = '_'
	}
	fmt.Printf("A palavra tem %d letras\n", numero)
	fmt.Println(string(progresso))
	fmt.Print("A letra e: ")
	fmt.Scan(escola_palavra)

}

func main() {
	//Player()
	Jogo_da_forca()
}
