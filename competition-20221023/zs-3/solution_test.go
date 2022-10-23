package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		cost []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{1, 3, 5, 2}, []int{2, 3, 1, 14}, 8},
		{[]int{2, 2, 2, 2, 2}, []int{4, 2, 8, 1, 3}, 0},
	}

	for i, v := range testGroup {
		result := minCost(v.nums, v.cost)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
