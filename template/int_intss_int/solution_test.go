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
		{6, [][]int{{0,1},{0,2},{1,2},{3,4}}, 3},
		{6, [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}}, 1},
	}

	for i, v := range testGroup {
		result := countCompleteComponents(v.n, DeepCopyIntss(v.edges))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
