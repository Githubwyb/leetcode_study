package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want int
	}

	testGroup := []testCase{
		{[][]int{{23, 99}, {83, 5}, {21, 76}, {34, 99}, {63, 23}}, 133},
		{[][]int{{1, 2, 4}, {3, 3, 1}}, 8},
		{[][]int{{10}}, 10},
	}

	for i, v := range testGroup {
		result := deleteGreatestValue(v.grid)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
