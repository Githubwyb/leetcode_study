package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want int64
	}

	testGroup := []testCase{
		{"12936", 11},
		{"5701283", 18},
		{"1010101010", 25},
	}

	for i, v := range testGroup {
		result := countSubstrings(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
