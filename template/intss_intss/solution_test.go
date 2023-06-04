package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want [][]int
	}

	testGroup := []testCase{
		{[][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}, [][]int{{1, 1, 0}, {1, 0, 1}, {0, 1, 1}}},
	}

	for i, v := range testGroup {
		result := differenceOfDistinctValues(DeepCopyIntss(v.grid))
		if !CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
