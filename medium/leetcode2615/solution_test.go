package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want []int64
	}

	testGroup := []testCase{
		{[]int{0, 5, 3, 1, 2, 8, 6, 6, 6}, []int64{0, 0, 0, 0, 0, 0, 3, 2, 3}},
		{[]int{1, 3, 1, 1, 2}, []int64{5, 0, 3, 4, 0}},
		{[]int{0, 5, 3}, []int64{0, 0, 0}},
	}

	for i, v := range testGroup {
		result := distance(common.DeepCopy(v.nums))
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := distance1(common.DeepCopy(v.nums))
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := distance2(common.DeepCopy(v.nums))
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
