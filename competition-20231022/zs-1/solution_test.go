package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{8, 6, 1, 5, 3}, 9},
		{[]int{5, 4, 8, 7, 10, 2}, 13},
		{[]int{6, 5, 4, 3, 4, 5}, -1},
	}

	for i, v := range testGroup {
		result := minimumSum(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
