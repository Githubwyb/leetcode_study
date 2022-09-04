package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		startPos int
		endPos   int
		k        int
		Want     int
	}

	testGroup := []testCase{
		{1, 2, 3, 3},
		{2, 5, 10, 0},
		{1, 1000, 999, 1},
		{989, 1000, 99, 934081896},
		{272, 270, 6, 15},
	}

	for i, v := range testGroup {
		result := numberOfWays(v.startPos, v.endPos, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
