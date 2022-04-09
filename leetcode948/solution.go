package main

import (
	_ "fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) (score int) {
	sort.Ints(tokens) // Ascending

	l := 0
	r := len(tokens) - 1
	point := 0
	// token isn't used up and point can pay for one token or power can be used for one token
	for l <= r && (power > tokens[l] || point > 0) {
		if power >= tokens[l] {
			power -= tokens[l]
			l++
			point++
			continue
		}

		// can't pay at least one tokens
		if point < score {
			break
		}
		score = point
		power += tokens[r]
		r--
		point--
	}
	if point > score {
		return point
	}
	return
}

func bagOfTokensScore1(tokens []int, power int) (score int) {
	sort.Ints(tokens) // Ascending

	l := 0
	r := len(tokens) - 1
	point := 0
	for l <= r && (power >= tokens[l] || point > 0) {
		for l <= r && power >= tokens[l] {
			power -= tokens[l]
			l++
			point++
		}

		if score < point {
			score = point
		}
		if l <= r && point > 0 {
			power += tokens[r]
			r--
			point--
		}
	}
	return
}
