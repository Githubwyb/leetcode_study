package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		coins []int
		edges [][]int
		Want  int
	}

	testGroup := []testCase{
		{[]int{1, 0, 0, 0, 0, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}

	for i, v := range testGroup {
		result := collectTheCoins(common.DeepCopy(v.coins), common.DeepCopyIntss(v.edges))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
