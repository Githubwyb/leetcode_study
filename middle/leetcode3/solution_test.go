package leetcode

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		S    string
		Want int
	}

	testGroup := []*testCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{" ", 1},
	}

	for i, v := range testGroup {
		result := LengthOfLongestSubstring(v.S)
		if result != v.Want {
			t.Fatalf("%d, input %s, expect %d but %d", i, v.S, v.Want, result)
		}
		fmt.Println(i, result)
	}

	for i, v := range testGroup {
		result := LengthOfLongestSubstring1(v.S)
		if result != v.Want {
			t.Fatalf("%d, input %s, expect %d but %d", i, v.S, v.Want, result)
		}
		fmt.Println(i, result)
	}

	for i, v := range testGroup {
		result := LengthOfLongestSubstring2(v.S)
		if result != v.Want {
			t.Fatalf("%d, input %s, expect %d but %d", i, v.S, v.Want, result)
		}
		fmt.Println(i, result)
	}
}
