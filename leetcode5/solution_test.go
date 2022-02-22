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

	testGroup := []*testCase{
		{"babab", "babab"},
		{"cbbd", "bb"},
	}

	for i, v := range testGroup {
		result := longestPalindrome(v.s)
		if result != v.Want {
			t.Fatalf("%d, s '%v', expect '%v' but '%v'", i, v.s, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := longestPalindrome1(v.s)
		if result != v.Want {
			t.Fatalf("%d, s '%v', expect '%v' but '%v'", i, v.s, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
