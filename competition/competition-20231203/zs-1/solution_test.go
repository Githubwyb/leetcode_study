package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want []int
	}

	testGroup := []testCase{
		{[]int{2, 4, 4}, []int{}},
		{[]int{1, 4, 3, 8, 5}, []int{1, 3}},
	}

	for i, v := range testGroup {
		result := findPeaks(DeepCopy(v.nums))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
