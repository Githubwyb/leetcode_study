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
		{[]int{3, 2, 1, 4}, 2},
		{[]int{1, 2}, -1},
		{[]int{2, 1, 3}, 2},
	}

	for i, v := range testGroup {
		result := findNonMinOrMax(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
