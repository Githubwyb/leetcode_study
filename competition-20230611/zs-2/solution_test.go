package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want string
	}

	testGroup := []testCase{
		{"cbabc", "baabc"},
		{"aa", "az"},
		{"a", "z"},
		{"acbbc", "abaab"},
		{"leetcode", "kddsbncd"},
	}

	for i, v := range testGroup {
		result := smallestString(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
