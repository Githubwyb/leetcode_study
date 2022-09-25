package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want []int
	}

	testGroup := []testCase{
		{[]int{2, 1, 1, 1, 3, 4, 1}, 2, []int{2, 3}},
		{[]int{2, 1, 1, 2}, 2, []int{}},
		{[]int{253747, 459932, 263592, 354832, 60715, 408350, 959296}, 2, []int{3}},
	}

	for i, v := range testGroup {
		result := goodIndices(v.nums, v.k)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		for i := range result {
			if result[i] != v.Want[i] {
				t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
}
