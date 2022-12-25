package main

func closetTarget(words []string, target string, startIndex int) int {
	if words[startIndex] == target {
		return 0
	}
	l, r := startIndex, startIndex
	for {
		l--
		if l < 0 {
			l = len(words) - 1
		}
		r++
		if r == len(words) {
			r = 0
		}
		if words[l] == target {
			return (startIndex - l + len(words)) % len(words)
		}
		if words[r] == target {
			return (r - startIndex + len(words)) % len(words)
		}
		if r == l {
			return -1
		}
	}
}
