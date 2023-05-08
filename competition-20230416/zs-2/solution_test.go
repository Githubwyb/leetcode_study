package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums     []int
		divisors []int
		Want     int
	}

	testGroup := []testCase{
		{[]int{73, 13, 20, 6}, []int{56, 75, 83, 26, 24, 53, 56, 61}, 24},
		{[]int{4, 7, 9, 3, 9}, []int{5, 2, 3}, 3},
		{[]int{20, 14, 21, 10}, []int{5, 7, 5}, 5},
		{[]int{12}, []int{10, 16}, 10},
	}

	for i, v := range testGroup {
		result := maxDivScore(common.DeepCopy(v.nums), common.DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
