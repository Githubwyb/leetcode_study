package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums1   []int
		nums2   []int
		queries [][]int
		Want    []int
	}

	testGroup := []testCase{
		{[]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}, []int{6, 10, 7}},
	}

	for i, v := range testGroup {
		result := maximumSumQueries(DeepCopy(v.nums1), DeepCopy(v.nums2), DeepCopyIntss(v.queries))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
