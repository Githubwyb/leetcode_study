package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{-2, -1, -1, 1, 2, 3}, 3},
		{[]int{-3, -2, -1, 0, 0, 1, 2}, 3},
		{[]int{5, 20, 66, 1314}, 4},
	}

	for i, v := range testGroup {
		result := maximumCount(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
