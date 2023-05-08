package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		in   []string
		Want []string
	}

	testGroup := []testCase{
		{[]string{"a", "b", "ab"}, []string{"ab"}},
		{[]string{"apple", "apples", "banana", "bananas", "bananaappleapples"}, []string{"bananaappleapples"}},
	}

	for i, v := range testGroup {
		result := ComboWords(v.in)
		if !CompareSlice(v.Want, result) {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
