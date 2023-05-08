package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums1 [][]int
		nums2 [][]int
		Want  [][]int
	}

	testGroup := []testCase{
		{[][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}, [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}}},
		{[][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}, [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}}},
	}

	for i, v := range testGroup {
		result := mergeArrays(common.DeepCopyIntss(v.nums1), common.DeepCopyIntss(v.nums2))
		if !common.CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := mergeArrays1(common.DeepCopyIntss(v.nums1), common.DeepCopyIntss(v.nums2))
		if !common.CompareSlices(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
