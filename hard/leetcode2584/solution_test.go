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
		{[]int{4, 7, 8, 15, 3, 5}, 2},
		{[]int{4, 7, 15, 8, 3, 5}, -1},
	}

	for i, v := range testGroup {
		result := findValidSplit(common.DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
