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
		{[][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}, 4, 7},
	}

	for i, v := range testGroup {
		result := countSubmatrices(DeepCopyIntss(v.grid), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
