package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want bool
	}

	testGroup := []testCase{
		{[][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}, true},
		{[][]int{{0, 3, 6}, {5, 8, 1}, {2, 7, 4}}, false},
	}

	for i, v := range testGroup {
		result := checkValidGrid(common.DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
