package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		op1  int
		op2  int
		Want int
	}

	testGroup := []testCase{
		{[]int{9}, 5, 1, 1, 0},
		{[]int{3, 10}, 5, 0, 1, 8},
		{[]int{275, 208, 626}, 553, 3, 2, 279},
		{[]int{10}, 3, 1, 1, 2},
		{[]int{2, 4, 3}, 3, 2, 1, 3},
		{[]int{2, 8, 3, 19, 3}, 3, 1, 1, 23},
	}

	for i, v := range testGroup {
		result := minArraySum(DeepCopy(v.nums), v.k, v.op1, v.op2)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
