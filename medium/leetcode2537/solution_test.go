package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int64
	}

	testGroup := []testCase{
		{[]int{1, 1, 1, 1, 1}, 10, 1},
		{[]int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
	}

	for i, v := range testGroup {
		result := countGood(v.nums, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
