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
		Want  int64
	}

	testGroup := []testCase{
		{3, [][]int{{0,0,1},{1,2,2},{0,2,3},{1,0,4}}, 23},
		{3, [][]int{{0,0,4},{0,1,2},{1,0,1},{0,2,3},{1,2,1}}, 17},
	}

	for i, v := range testGroup {
		result := matrixSumQueries(v.n, DeepCopyIntss(v.edges))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
