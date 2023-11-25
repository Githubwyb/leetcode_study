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
		{[]int{1,1,1,3,3,3,1,2,1,1,1,2,1}, 5},
		{[]int{1,1,2,1,1,1,3,1,2,3}, 4},
		{[]int{3, 2, 3, 2, 3}, 2},
		{[]int{10, 10, 10, 3, 1, 1}, 4},
	}

	for i, v := range testGroup {
		result := minGroupsForValidAssignment(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
