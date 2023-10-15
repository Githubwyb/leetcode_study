package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		Want [][]int
	}

	testGroup := []testCase{
		{10, [][]int{{3, 7}, {5, 5}}},
	}

	for i, v := range testGroup {
		result := findPrimePairs(v.n)
		if !CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
