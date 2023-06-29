package main

import (
	"fmt"
	"strings"
)

func countLetters(frase string) map[string]map[rune]int {
	palavras := strings.Fields(frase)
	resultado := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		letras := make(map[rune]int)
		for _, letra := range palavra {
			letras[letra]++
		}
		resultado[palavra] = letras
	}

	return resultado
}

func main() {
	frase := "A frase de exemplo"
	resultado := countLetters(frase)

	for palavra, letras := range resultado {
		fmt.Println("Palavra:", palavra)
		for letra, count := range letras {
			fmt.Printf("Letra: %c, Contagem: %d\n", letra, count)
		}
		fmt.Println()
	}
}
