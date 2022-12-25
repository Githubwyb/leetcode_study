package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n     int
		roads [][]int
		Want  int
	}

	testGroup := []testCase{
		{4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}, 5},
		{4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}, 2},
	}

	for i, v := range testGroup {
		result := minScore(v.n, v.roads)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
