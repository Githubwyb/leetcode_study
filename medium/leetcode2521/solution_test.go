package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{2, 4, 3, 7, 10, 6}, 4},
		{[]int{2, 4, 8, 16}, 1},
	}

	for i, v := range testGroup {
		result := distinctPrimeFactors(common.DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
