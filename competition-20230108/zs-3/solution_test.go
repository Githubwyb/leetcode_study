package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		word1 string
		word2 string
		Want  bool
	}

	testGroup := []testCase{
		{"az", "a", true},
		{"aab", "bccd", true},
		{"a", "bb", false},
		{"ac", "b", false},
		{"abcc", "aab", true},
		{"abcde", "fghij", true},
	}

	for i, v := range testGroup {
		result := isItPossible(v.word1, v.word2)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
