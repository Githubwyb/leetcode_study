package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		word string
		k    int
		Want int
	}

	testGroup := []testCase{
		{"aabcaba", 0, 3},
		{"dabdcbdcdcd", 2, 2},
		{"aaabaaa", 2, 1},
	}

	for i, v := range testGroup {
		result := minimumDeletions(v.word, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
