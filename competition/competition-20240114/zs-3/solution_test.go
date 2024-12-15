package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int64
		time int
		Want int64
	}

	testGroup := []testCase{
		{9, 1, 6},
		{7, 2, 9},
	}

	for i, v := range testGroup {
		result := findMaximumNumber(v.n, v.time)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
