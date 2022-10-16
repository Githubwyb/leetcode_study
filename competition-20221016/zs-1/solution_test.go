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
		{[]int{-1, 2, -3, 3}, 3},
		{[]int{-1, 10, 6, 7, -7, 1}, 7},
		{[]int{-10, 8, 6, 7, -2, -3}, -1},
	}

	for i, v := range testGroup {
		result := findMaxK(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
