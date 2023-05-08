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
		{[]int{10, 4, 8, 3}, []int64{15, 1, 11, 22}},
	}

	for i, v := range testGroup {
		result := distance(common.DeepCopy(v.nums))
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
