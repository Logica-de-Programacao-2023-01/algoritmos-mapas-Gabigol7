package main

import "fmt"

func fibonacciSequence(n int) map[int]int {
	seq := make(map[int]int)

	seq[0] = 0
	seq[1] = 1

	for i := 2; ; i++ {
		seq[i] = seq[i-1] + seq[i-2]
		if seq[i] > n {
			break
		}
	}

	return seq
}

func main() {
	n := 100 // O número máximo até o qual deseja gerar a sequência de Fibonacci

	seq := fibonacciSequence(n)

	for index, value := range seq {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
