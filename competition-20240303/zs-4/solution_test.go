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
		{[]int{2, 47, 3, 3, 65, 7, 97, 4}, []int{2, 65, 7, 97, 4, 47, 3, 3}},
		{[]int{2, 1, 3, 3}, []int{2, 3, 1, 3}},
		{[]int{5, 14, 3, 1, 2}, []int{5, 3, 1, 2, 14}},
		{[]int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
	}

	for i, v := range testGroup {
		result := resultArray(DeepCopy(v.nums))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
