package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		cost       []int
		k          int
		candidates int
		Want       int64
	}

	testGroup := []testCase{
		{[]int{31,25,72,79,74,65,84,91,18,59,27,9,81,33,17,58}, 11, 2, 423},
		{[]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			7,
			3, 13},
		{[]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4, 11},
		{[]int{1, 2, 4, 1}, 3, 3, 4},
	}

	for i, v := range testGroup {
		result := totalCost(common.DeepCopy(v.cost), v.k, v.candidates)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := totalCost1(common.DeepCopy(v.cost), v.k, v.candidates)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
