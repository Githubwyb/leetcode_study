package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		k    int
		in   []int
		Want int
	}

	testGroup := []testCase{
		{3, []int{2, 2, 2, 2}, 4},
		{1, []int{1, 2, 3}, 3},
		{1, []int{1}, 1},
	}

	for i, v := range testGroup {
		result := applesSum(v.k, v.in)
		if result != v.Want {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
