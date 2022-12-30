package main

import "math"

func closetTarget(words []string, target string, startIndex int) int {
	if words[startIndex] == target {
		return 0
	}
	l, r := startIndex, startIndex
	n := 0
	for {
		l--
		if l < 0 {
			l = len(words) - 1
		}
		r++
		if r == len(words) {
			r = 0
		}
		n++
		if words[l] == target || words[r] == target {
			return n
		}
		if r == l {
			return -1
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func closetTarget1(words []string, target string, startIndex int) int {
	result := math.MaxInt
	for i := range words {
		if words[i] != target {
			continue
		}
		d := abs(startIndex - i)
		d = min(d, len(words)-d)
		result = min(result, d)
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}
