package main

func divisibilityArray(word string, m int) []int {
	result := make([]int, len(word))
	n := 0
	for i, v := range word {
		n = n*10 + int(v-'0')
		n %= m
		if n == 0 {
			result[i] = 1
		}
	}
	return result
}
