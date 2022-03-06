package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		p    string
		Want bool
	}

	testGroup := []*testCase{
		{"aa", "a*", true},
		{"ab", ".*c", false},
		{"aaca", "ab*a*c*a", true},
		{"aaa", "ab*a", false},
		{"abcd", "d*", false},
		{"aa", "a", false},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"aaa", "ab*a*c*a", true},
	}

	for i, v := range testGroup {
		result := isMatch(v.s, v.p)
		if result != v.Want {
			t.Fatalf("%d, s %v p %v, expect '%v' but '%v'", i, v.s, v.p, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := isMatch1(v.s, v.p)
		if result != v.Want {
			t.Fatalf("%d, s %v p %v, expect '%v' but '%v'", i, v.s, v.p, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
