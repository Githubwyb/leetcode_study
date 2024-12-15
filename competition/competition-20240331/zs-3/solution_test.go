package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{0, 1, 1, 1}, 5},
		{[]int{1, 0, 1, 0}, 10},
	}

	for i, v := range testGroup {
		result := countAlternatingSubarrays(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
