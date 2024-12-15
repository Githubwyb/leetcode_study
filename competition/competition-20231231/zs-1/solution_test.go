package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want bool
	}

	testGroup := []testCase{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{2, 4, 8, 16}, true},
		{[]int{1, 3, 5, 7, 9}, false},
	}

	for i, v := range testGroup {
		result := hasTrailingZeros(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
