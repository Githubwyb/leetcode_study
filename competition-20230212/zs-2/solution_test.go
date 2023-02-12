package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums  []int
		lower int
		upper int
		Want  int64
	}

	testGroup := []testCase{
		{[]int{0, 0, 0, 0, 0, 0}, 0, 0, 15},
		{[]int{0, 1, 7, 4, 4, 5}, 3, 6, 6},
		{[]int{1, 7, 9, 2, 5}, 11, 11, 1},
	}

	for i, v := range testGroup {
		result := countFairPairs(common.DeepCopy(v.nums), v.lower, v.upper)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
