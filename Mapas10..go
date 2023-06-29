package main

import (
	"fmt"
)

func countPairs(slice []int) map[[2]int]int {
	pairCounts := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			pairCounts[pair]++
		}
	}

	return pairCounts
}

func main() {
	slice := []int{1, 2, 3, 2, 4, 1, 3, 1}
	pairCounts := countPairs(slice)

	for pair, count := range pairCounts {
		fmt.Printf("%v: %d\n", pair, count)
	}
}
