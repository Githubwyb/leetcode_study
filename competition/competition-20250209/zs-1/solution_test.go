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
		{ // 案例1
			grid: [][]int{
				{1, 7, 3},
				{9, 8, 2},
				{4, 5, 6},
			},
			Want: [][]int{
				{8, 2, 3},
				{9, 6, 7},
				{4, 5, 1},
			},
		},
		{ // 案例2
			grid: [][]int{{0, 1}, {1, 2}},
			Want: [][]int{{2, 1}, {1, 0}},
		},
		{ // 案例3
			grid: [][]int{{1}},
			Want: [][]int{{1}},
		},
	}

	for i, v := range testGroup {
		result := sortMatrix(DeepCopyIntss(v.grid))
		if !CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
