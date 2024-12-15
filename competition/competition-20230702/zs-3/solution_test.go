package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{5, 4, 2, 4}, 8},
		{[]int{1, 2, 3}, 6},
	}

	for i, v := range testGroup {
		result := continuousSubarrays(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
