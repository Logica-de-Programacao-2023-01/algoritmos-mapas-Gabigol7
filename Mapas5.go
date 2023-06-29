package main

func countCharacters(str string) map[rune]int {
	freqMap := make(map[rune]int)

	for _, char := range str {
		freqMap[char]++
	}

	return freqMap
}
