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
		{[][]int{{0,1},{0,0}}, 0},
		{[][]int{{0,0,1},{1,0,1},{0,0,0}}, 1},
	}

	for i, v := range testGroup {
		result := findChampion(DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
