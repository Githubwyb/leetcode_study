package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		tree    []interface{}
		queries []int
		Want    [][]int
	}

	testGroup := []testCase{
		{[]interface{}{6, 2, 13, 1, 4, 9, 15, nil, nil, nil, nil, nil, nil, 14}, []int{2, 5, 16}, [][]int{{2, 2}, {4, 6}, {15, -1}}},
		{[]interface{}{4, nil, 9}, []int{3}, [][]int{{-1, 4}}},
	}

	for i, v := range testGroup {
		result := closestNodes(common.MakeTreeNode(v.tree), v.queries)
		if !common.CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
