package main

import "sort"

func sortNormalChar(input string) string {
	charArr := make([]int, 0, len(input))
	indexArr := make([]int, 0, len(input))
	for i, v := range input {
		if v > 'z' || v < 'a' {
			continue
		}
		charArr = append(charArr, int(v))
		indexArr = append(indexArr, i)
	}
	sort.Ints(charArr)
	outputBytes := []byte(input)
	for i, v := range indexArr {
		outputBytes[v] = byte(charArr[i])
	}
	return string(outputBytes)
}
