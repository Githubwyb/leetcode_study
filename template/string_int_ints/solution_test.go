package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		word string
		m    int
		Want []int
	}

	testGroup := []testCase{
		{"998244353", 3, []int{15, 1, 11, 22}},
	}

	for i, v := range testGroup {
		result := divisibilityArray(v.word, v.m)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
