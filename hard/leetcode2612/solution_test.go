package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n      int
		p      int
		banned []int
		k      int
		Want   []int
	}

	testGroup := []testCase{
		{5, 3, []int{0}, 3, []int{-1, 1, -1, 0, -1}},
		{3, 2, []int{}, 2, []int{2, 1, 0}},
		{5, 0, []int{}, 2, []int{0, 1, 2, 3, 4}},
		{4, 0, []int{1, 2}, 4, []int{0, -1, -1, 1}},
		{5, 0, []int{2, 4}, 3, []int{0, -1, -1, -1, -1}},
		{4, 2, []int{0, 1, 3}, 1, []int{-1, -1, 0, -1}},
	}

	for i, v := range testGroup {
		result := minReverseOperations(v.n, v.p, common.DeepCopy(v.banned), v.k)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
