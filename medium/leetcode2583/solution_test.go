package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		root []interface{}
		k    int
		Want int64
	}

	testGroup := []testCase{
		{[]interface{}{5, 8, 9, 2, 1, 3, 7, 4, 6}, 2, 13},
		{[]interface{}{1, 2, nil, 3}, 1, 3},
	}

	for i, v := range testGroup {
		result := kthLargestLevelSum(MakeTreeNode(v.root), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
