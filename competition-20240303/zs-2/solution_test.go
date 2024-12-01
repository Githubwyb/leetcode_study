package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		k    int
		Want int
	}

	testGroup := []testCase{
		{[][]int{{7, 6, 3}, {6, 6, 1}}, 18, 4},
		{[][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20, 6},
	}

	for i, v := range testGroup {
		result := countSubmatrices(DeepCopyIntss(v.grid), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
