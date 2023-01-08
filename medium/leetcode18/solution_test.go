package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		Want   [][]int
	}

	testGroup := []*testCase{
		{[]int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}

	for i, v := range testGroup {
		result := fourSum(common.DeepCopy(v.nums), v.target)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
