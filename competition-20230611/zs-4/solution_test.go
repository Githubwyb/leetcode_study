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
		{[]int{28, 18, 97, 6, 14, 27}, []int{19, 11, 1, 23, 7, 52}, [][]int{{91, 29}, {33, 60}, {80, 54}, {24, 19}}, []int{-1, -1, -1, 79}},
		{[]int{48, 98, 61, 28, 14, 7}, []int{36, 12, 33, 67, 14, 73}, [][]int{{98, 52}, {8, 42}}, []int{-1, 95}},
		{[]int{23}, []int{38}, [][]int{{68, 22}}, []int{-1}},
		{[]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}, []int{6, 10, 7}},
		{[]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}, []int{9, 9, 9}},
		{[]int{2, 1}, []int{2, 3}, [][]int{{3, 3}}, []int{-1}},
	}

	for i, v := range testGroup {
		result := maximumSumQueries(DeepCopy(v.nums1), DeepCopy(v.nums2), DeepCopyIntss(v.queries))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
