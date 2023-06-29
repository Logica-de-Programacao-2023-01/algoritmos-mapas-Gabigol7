package main

import (
	"fmt"
)

func main() {
	// Exemplo de entrada com uma lista de mapas
	wordCounts := []map[string]int{
		{"apple": 2, "banana": 3, "orange": 1},
		{"apple": 1, "orange": 2, "pear": 4},
		{"banana": 5, "orange": 3, "grape": 2},
	}

	// Chamada da função para obter o mapa com a soma das contagens
	totalCounts := sumWordCounts(wordCounts)

	// Imprimir o resultado
	fmt.Println(totalCounts)
}

func sumWordCounts(wordCounts []map[string]int) map[string]int {
	totalCounts := make(map[string]int)

	for _, counts := range wordCounts {
		for word, count := range counts {
			totalCounts[word] += count
		}
	}

	return totalCounts
}
