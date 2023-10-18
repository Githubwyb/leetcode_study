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
		{[][]int{{1, 2}, {3, 4}}, [][]int{{24, 12}, {8, 6}}},
		{[][]int{{12345}, {2}, {1}}, [][]int{{2}, {0}, {0}}},
	}

	for i, v := range testGroup {
		result := constructProductMatrix(DeepCopyIntss(v.grid))
		if !CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
