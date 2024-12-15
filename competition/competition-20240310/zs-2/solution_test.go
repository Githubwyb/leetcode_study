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
		Want int64
	}

	testGroup := []testCase{
		{[]int{1, 2, 3}, 2, 4},
		{[]int{1, 1, 1, 1}, 2, 1},
		{[]int{2, 3, 4, 5}, 1, 5},
	}

	for i, v := range testGroup {
		result := maximumHappinessSum(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
