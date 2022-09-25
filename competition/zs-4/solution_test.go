package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		vals  []int
		edges [][]int
		Want  int
	}

	testGroup := []testCase{
		{[]int{1, 3, 2, 1, 3}, [][]int{[]int{0, 1}, []int{0, 2}, []int{2, 3}, []int{2, 4}}, 6},
	}

	for i, v := range testGroup {
		result := numberOfGoodPaths(v.vals, v.edges)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
