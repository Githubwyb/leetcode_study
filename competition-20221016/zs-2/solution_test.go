package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{1, 13, 10, 12, 31}, 6},
		{[]int{2, 2, 2}, 1},
	}

	for i, v := range testGroup {
		result := countDistinctIntegers(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
