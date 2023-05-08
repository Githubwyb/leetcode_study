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
		{"01000111", 6},
		{"00111", 4},
		{"111", 0},
	}

	for i, v := range testGroup {
		result := findTheLongestBalancedSubstring(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
