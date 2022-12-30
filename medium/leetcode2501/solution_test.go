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
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{4, 3, 6, 16, 8, 2}, 3},
		{[]int{2, 3, 5, 6, 7}, -1},
		{[]int{2, 2}, -1},
	}

	for i, v := range testGroup {
		result := longestSquareStreak(deepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := longestSquareStreak1(deepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
