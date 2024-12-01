package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want int
	}

	testGroup := []testCase{
		{"aaaa", 2},
		{"abcdef", -1},
		{"abcaba", 1},
	}

	for i, v := range testGroup {
		result := maximumLength(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
