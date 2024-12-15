package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		arr  []string
		Want []string
	}

	testGroup := []testCase{
		{[]string{"cab", "ad", "bad", "c"}, []string{"ab", "", "ba", ""}},
		{[]string{"abc", "bcd", "abcd"}, []string{"", "", "abcd"}},
	}

	for i, v := range testGroup {
		result := shortestSubstrings(DeepCopy(v.arr))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
