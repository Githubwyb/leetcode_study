package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums1 []int
		nums2 []int
		Want  []int
	}

	testGroup := []testCase{
		{[]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, []int{6, 10, 7}},
	}

	for i, v := range testGroup {
		result := assignElements(DeepCopy(v.nums1), DeepCopy(v.nums2))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
