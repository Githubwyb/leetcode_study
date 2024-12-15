package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		l    int
		r    int
		Want int
	}

	testGroup := []testCase{
		{[]int{3, -2, 1, 4}, 2, 3, 1},
		{[]int{-2, 2, -3, 1}, 2, 3, -1},
		{[]int{1, 2, 3, 4}, 2, 4, 3},
	}

	for i, v := range testGroup {
		result := minimumSumSubarray(DeepCopy(v.nums), v.l, v.r)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
