package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		edges1 [][]int
		edges2 [][]int
		k      int
		Want   []int
	}

	testGroup := []testCase{
		{[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}, 2, []int{9, 7, 9, 8, 8}},
	}

	for i, v := range testGroup {
		result := maxTargetNodes(DeepCopyIntss(v.edges1), DeepCopyIntss(v.edges2), v.k)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
