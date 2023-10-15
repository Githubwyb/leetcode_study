package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{5, 3, 4, 1, 1}, 13},
		{[]int{6, 5, 3, 9, 2, 7}, 22},
		{[]int{3, 2, 5, 5, 2, 3}, 18},
	}

	for i, v := range testGroup {
		result := maximumSumOfHeights(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
