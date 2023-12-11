package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		mat    [][]int
		target int
		Want   []int
	}

	testGroup := []testCase{
		{[][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}, 2, []int{0, 2}},
		{[][]int{{39, 3, 1000, 1000}}, 17, []int{}},
	}

	for i, v := range testGroup {
		result := getGoodIndices(DeepCopyIntss(v.mat), v.target)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
