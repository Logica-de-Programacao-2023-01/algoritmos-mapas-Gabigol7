package main

import (
	"fmt"
	"strings"
)

func countWords(str string) map[string]int {
	words := strings.Fields(str) // Divide a string em palavras
	wordCount := make(map[string]int)

	// Conta a ocorrência de cada palavra
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Olá mundo olá Go mundo Go"
	wordCount := countWords(text)

	// Exibe a contagem de palavras
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
