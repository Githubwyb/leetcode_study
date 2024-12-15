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
		{[]int{1, 3, 2, 3, 3}, 2, 6},
		{[]int{1, 4, 2, 1}, 3, 0},
	}

	for i, v := range testGroup {
		result := countSubarrays(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
