package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums            []int
		indexDifference int
		valueDifference int
		Want            []int
	}

	testGroup := []testCase{
		{[]int{5, 1, 4, 1}, 2, 4, []int{0, 3}},
		{[]int{2, 1}, 0, 0, []int{0, 0}},
		{[]int{1, 2, 3}, 2, 4, []int{-1, -1}},
	}

	for i, v := range testGroup {
		result := findIndices(DeepCopy(v.nums), v.indexDifference, v.valueDifference)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
