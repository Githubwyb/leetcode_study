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
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{5, 4, 3, 8}, []int{5, 3, 4, 8}},
	}

	for i, v := range testGroup {
		result := resultArray(DeepCopy(v.nums))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
