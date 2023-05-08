package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		t    string
		Want int
	}

	testGroup := []testCase{
		{"coaching", "coding", 4},
		{"abcde", "a", 0},
		{"z", "abcde", 5},
	}

	for i, v := range testGroup {
		result := appendCharacters(v.s, v.t)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
