package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n       int
		queries [][]int
		Want    [][]int
	}

	testGroup := []testCase{
		{
			13,
			[][]int{{3, 1, 7, 3}, {7, 5, 7, 8}, {4, 12, 6, 12}, {2, 8, 6, 11}, {9, 11, 10, 11}, {9, 3, 11, 11}, {0, 12, 10, 12}, {10, 5, 11, 12}, {4, 7, 6, 12}, {0, 2, 9, 6}, {12, 7, 12, 11}, {2, 7, 3, 8}, {2, 9, 6, 12}, {10, 7, 10, 12}, {11, 6, 11, 7}, {3, 2, 12, 9}},
			[][]int{{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, {0, 1, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 2}, {0, 1, 3, 3, 2, 2, 2, 2, 3, 4, 3, 3, 4}, {0, 1, 3, 3, 2, 2, 2, 2, 3, 4, 3, 3, 4}, {0, 1, 3, 3, 2, 2, 2, 2, 3, 4, 3, 3, 4}, {0, 1, 3, 3, 2, 3, 3, 2, 2, 1, 0, 0, 1}, {0, 0, 2, 2, 2, 2, 2, 1, 1, 1, 0, 0, 1}, {0, 0, 2, 3, 3, 3, 3, 2, 2, 2, 1, 2, 1}, {0, 0, 1, 2, 2, 3, 3, 4, 4, 4, 3, 4, 3}, {0, 0, 1, 2, 2, 3, 4, 4, 3, 3, 2, 2, 1}, {0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 0}},
		},
		{3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}, [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}}},
		{2, [][]int{{0, 0, 1, 1}}, [][]int{{1, 1}, {1, 1}}},
	}

	for i, v := range testGroup {
		result := rangeAddQueries(v.n, v.queries)
		if !common.CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := rangeAddQueries1(v.n, v.queries)
		if !common.CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
