package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{6, 14, 20}, 840},
		{[]int{2, 4, 8, 16}, 64},
		{[]int{1, 2, 3, 4, 5}, 60},
		{[]int{3}, 9},
	}

	for i, v := range testGroup {
		result := maxScore(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}

	for i, v := range testGroup {
		result := maxScore1(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
