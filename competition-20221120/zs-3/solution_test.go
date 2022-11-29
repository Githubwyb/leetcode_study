package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		roads [][]int
		seats int
		Want  int64
	}

	testGroup := []testCase{
		{[][]int{{0, 1}, {0, 2}, {0, 3}}, 5, 3},
		{[][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2, 7},
		{[][]int{}, 1, 0},
	}

	for i, v := range testGroup {
		result := minimumFuelCost(v.roads, v.seats)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
