package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n     int
		edges [][]int
		Want  int64
	}

	testGroup := []testCase{
		{10, [][]int{{10, 9}, {2, 10}, {1, 10}, {3, 2}, {6, 10}, {4, 3}, {8, 6}, {5, 8}, {7, 6}}, 16},
		{5, [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}}, 4},
		{6, [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {3, 6}}, 6},
	}

	for i, v := range testGroup {
		result := countPaths(v.n, DeepCopyIntss(v.edges))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
