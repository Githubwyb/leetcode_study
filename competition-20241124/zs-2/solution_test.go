package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		t    string
		k    int
		Want bool
	}

	testGroup := []testCase{
		{"abcd", "cdab", 2, true},
		{"aabbcc", "bbaacc", 3, true},
		{"aabbcc", "bbaacc", 2, false},
	}

	for i, v := range testGroup {
		result := isPossibleToRearrange(v.s, v.t, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
