package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		k    int
		Want string
	}

	testGroup := []testCase{
		{"100011001", 3, "11001"},
	}

	for i, v := range testGroup {
		result := shortestBeautifulSubstring(v.s, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
