package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		c    byte
		Want int64
	}

	testGroup := []testCase{
		{"abada", 'a', 6},
	}

	for i, v := range testGroup {
		result := countSubstrings(v.s, v.c)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
