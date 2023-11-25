package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		edges [][]int
		coins []int
		Want  int64
	}

	testGroup := []testCase{
		{[][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {4, 5}}, []int{5, 2, 5, 2, 1, 1}, 2},
		{[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, []int{20, 10, 9, 7, 4, 3, 5}, 2},
	}

	for i, v := range testGroup {
		result := maximumScoreAfterOperations(DeepCopyIntss(v.edges), DeepCopy(v.coins))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
