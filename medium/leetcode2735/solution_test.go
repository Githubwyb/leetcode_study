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
		{[]int{20, 1, 15}, 5, 13},
		{[]int{1, 2, 3}, 4, 6},
	}

	for i, v := range testGroup {
		result := minCost(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
