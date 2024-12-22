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
		{"0110", 2, 1},
		{"000", 0, 3},
		{"000001", 1, 2},
		{"0000", 2, 1},
		{"0101", 0, 1},
	}

	for i, v := range testGroup {
		result := minLength(v.word, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
