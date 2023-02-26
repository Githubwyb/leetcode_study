package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{3, 5, 2, 4}, 2},
		{[]int{9, 2, 5, 4}, 4},
		{[]int{7, 6, 8}, 0},
	}

	for i, v := range testGroup {
		result := maxNumOfMarkedIndices(common.DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
