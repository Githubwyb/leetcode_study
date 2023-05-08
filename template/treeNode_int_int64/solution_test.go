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
		{[]interface{}{7, 52, 2, 4}, 2, 13},
	}

	for i, v := range testGroup {
		result := kthLargestLevelSum(MakeTreeNode(v.root), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
