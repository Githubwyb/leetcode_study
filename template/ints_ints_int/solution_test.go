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
		{[]int{1, 0, 0, 0, 0, 1}, []int{2}, 1},
	}

	for i, v := range testGroup {
		result := maxDivScore(common.DeepCopy(v.nums), common.DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
