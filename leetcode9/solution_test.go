package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		x    int
		Want bool
	}

	testGroup := []*testCase{
		{121, true},
		{-232, false},
		{10, false},
	}

	for i, v := range testGroup {
		result := isPalindrome(v.x)
		if result != v.Want {
			t.Fatalf("%d, x %v expect '%v' but '%v'", i, v.x, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := isPalindrome1(v.x)
		if result != v.Want {
			t.Fatalf("%d, x %v expect '%v' but '%v'", i, v.x, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
