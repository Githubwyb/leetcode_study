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
		{[]int{1, 15, 6, 3}, 9},
		{[]int{1, 2, 3, 4}, 0},
	}

	for i, v := range testGroup {
		result := differenceOfSum(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
