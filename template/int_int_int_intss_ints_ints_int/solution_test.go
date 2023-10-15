package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n           int
		k           int
		budget      int
		composition [][]int
		stock       []int
		cost        []int
		Want        int
	}

	testGroup := []testCase{
		{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 0}, []int{1, 2, 3}, 2},
		{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 100}, []int{1, 2, 3}, 5},
		{2, 3, 10, [][]int{{2, 1}, {1, 2}, {1, 1}}, []int{1, 1}, []int{5, 5}, 2},
	}

	for i, v := range testGroup {
		result := maxNumberOfAlloys(v.n, v.k, v.budget, DeepCopyIntss(v.composition), DeepCopy(v.stock), DeepCopy(v.cost))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
