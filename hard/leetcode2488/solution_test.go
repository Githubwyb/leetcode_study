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
		k    int
		Want int
	}

	testGroup := []testCase{
		{[]int{5, 19, 11, 15, 13, 16, 4, 6, 2, 7, 10, 8, 18, 20, 1, 3, 17, 9, 12, 14}, 6, 13},
		{[]int{4, 1, 3, 2}, 1, 3},
		{[]int{3, 2, 1, 4, 5}, 4, 3},
		{[]int{2, 3, 1}, 3, 1},
	}

	for i, v := range testGroup {
		result := countSubarrays(deepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := countSubarrays1(deepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
