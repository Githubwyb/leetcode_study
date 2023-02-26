package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want int
	}

	testGroup := []testCase{
		{[][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}, 7},
		{[][]int{{0, 2, 4}, {3, 2, 1}, {1, 0, 4}}, -1},
	}

	for i, v := range testGroup {
		result := minimumTime(common.DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := minimumTime1(common.DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
