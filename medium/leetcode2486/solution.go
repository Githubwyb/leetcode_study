package main

func appendCharacters(s string, t string) int {
	i := 0
	for _, v := range s {
		if byte(v) == t[i] {
			i++
			if i == len(t) {
				return 0
			}
		}
	}
	return len(t) - i
}
