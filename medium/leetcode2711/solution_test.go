package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want [][]int
	}

	testGroup := []testCase{
		{[][]int{{6, 28, 37, 34, 12, 30, 43, 35, 6}, {21, 47, 38, 14, 31, 49, 11, 14, 49}, {6, 12, 35, 17, 17, 2, 45, 27, 43}, {34, 41, 30, 28, 45, 24, 50, 20, 4}}, [][]int{{3, 3, 3, 3, 3, 3, 2, 1, 0}, {2, 1, 1, 1, 1, 1, 1, 0, 1}, {1, 0, 1, 1, 1, 1, 1, 1, 2}, {0, 1, 2, 3, 3, 3, 3, 3, 3}}},
		{[][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}, [][]int{{1, 1, 0}, {1, 0, 1}, {0, 1, 1}}},
		{[][]int{{1}}, [][]int{{0}}},
	}

	for i, v := range testGroup {
		result := differenceOfDistinctValues(DeepCopyIntss(v.grid))
		if !CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
