package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		arr  []int
		in   int
		Want int
	}

	testGroup := []testCase{
		{[]int{1, 2, 2, 3, 4}, 0, 0},
		{[]int{1, 2, 2, 3, 4}, 2, 3},
		{[]int{1, 2, 2, 3, 4}, 5, 5},
	}

	for i, v := range testGroup {
		result := getInsertLocation(v.arr, v.in)
		if result != v.Want {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
