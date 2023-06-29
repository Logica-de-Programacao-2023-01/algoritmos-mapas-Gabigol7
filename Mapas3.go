package main

import "fmt"

func somarValoresMapa(m map[string]int) int {
	soma := 0
	for _, valor := range m {
		soma += valor
	}
	return soma
}

func main() {
	// Exemplo de mapa com valores inteiros
	mapa := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	// Chamando a função para somar os valores do mapa
	resultado := somarValoresMapa(mapa)
	fmt.Println("Soma dos valores:", resultado)
}
