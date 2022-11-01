package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		minK int
		maxK int
		Want int64
	}

	testGroup := []testCase{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	}

	for i, v := range testGroup {
		result := countSubarrays(v.nums, v.minK, v.maxK)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
