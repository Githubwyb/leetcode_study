package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums  []int
		lower int
		upper int
		Want  int
	}

	testGroup := []testCase{
		{[]int{-2, 5, -1}, -2, 2, 3},
		{[]int{0}, 0, 0, 0},
	}

	for i, v := range testGroup {
		result := countRangeSum(DeepCopy(v.nums), v.lower, v.upper)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
