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
		{[]int{2, 1, 4, 3}, 2},
		{[]int{2, 4, 1, 3}, 3},
		{[]int{1, 3, 4, 2, 5}, 0},
	}

	for i, v := range testGroup {
		result := semiOrderedPermutation(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
