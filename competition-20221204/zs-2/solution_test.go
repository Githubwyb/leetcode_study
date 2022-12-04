package main

import (
	"fmt"
	"testing"
)

func deepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}

func TestSolution(t *testing.T) {
	type testCase struct {
		skill []int
		Want  int64
	}

	testGroup := []testCase{
		{[]int{3, 2, 5, 1, 3, 4}, 22},
		{[]int{3, 4}, 12},
		{[]int{1, 1, 2, 3}, -1},
	}

	for i, v := range testGroup {
		result := dividePlayers(deepCopy(v.skill))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
