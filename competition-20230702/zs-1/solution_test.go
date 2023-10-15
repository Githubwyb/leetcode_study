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
		Want int
	}

	testGroup := []testCase{
		{[]int{3, 2, 5, 4}, 5, 3},
		{[]int{1, 2}, 2, 1},
		{[]int{2, 3, 4, 5}, 4, 3},
	}

	for i, v := range testGroup {
		result := longestAlternatingSubarray(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
