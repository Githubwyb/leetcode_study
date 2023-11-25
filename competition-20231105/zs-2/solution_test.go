package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n     int
		edges [][]int
		Want  int
	}

	testGroup := []testCase{
		{3, [][]int{{0,1},{1,2}}, 0},
		{4, [][]int{{0,2},{1,3},{1,2}}, -1},
	}

	for i, v := range testGroup {
		result := findChampion(v.n, DeepCopyIntss(v.edges))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
