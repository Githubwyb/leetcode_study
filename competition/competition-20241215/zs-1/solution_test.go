package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want int
	}

	testGroup := []testCase{
		{[][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}}, 1},
		{[][]int{{10, 5}, {1, 7}}, 10},
	}

	for i, v := range testGroup {
		result := buttonWithLongestTime(DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
