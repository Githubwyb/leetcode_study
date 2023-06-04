package main

func minimizedStringLength(s string) int {
	sMap := make(map[byte]bool)
	for _, v := range s {
		sMap[byte(v)] = true
	}
	return len(sMap)
}
