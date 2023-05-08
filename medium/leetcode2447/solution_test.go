package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int
	}

	testGroup := []testCase{
		{[]int{9, 3, 1, 2, 6, 3}, 3, 4},
		{[]int{4}, 7, 0},
	}

	for i, v := range testGroup {
		result := subarrayGCD(v.nums, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
