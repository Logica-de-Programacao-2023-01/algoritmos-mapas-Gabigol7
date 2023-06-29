package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	// Converte a string em um slice de caracteres
	chars := strings.Split(s, "")
	// Ordena o slice de caracteres em ordem crescente
	sort.Strings(chars)
	// Junta os caracteres em uma única string novamente
	return strings.Join(chars, "")
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		// Ordena as letras da palavra em ordem alfabética
		sortedWord := sortString(strings.ToLower(word))
		// Adiciona a palavra ao grupo de anagramas correspondente
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	return anagrams
}

func main() {
	// Exemplo de uso
	wordSlice := []string{"car", "arc", "dog", "god", "listen", "silent", "pot", "top"}
	anagramMap := findAnagrams(wordSlice)

	for key, value := range anagramMap {
		fmt.Printf("%s: %v\n", key, value)
	}
}
