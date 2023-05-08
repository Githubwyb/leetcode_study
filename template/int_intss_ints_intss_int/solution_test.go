package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
		Want  int
	}

	testGroup := []testCase{
		{4, [][]int{{0, 1}, {1, 2}, {1, 3}}, []int{2, 2, 10, 6}, [][]int{{0, 3}, {2, 1}, {2, 3}}, 23},
	}

	for i, v := range testGroup {
		result := minimumTotalPrice(v.n, common.DeepCopyIntss(v.edges), common.DeepCopy(v.price), common.DeepCopyIntss(v.trips))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
