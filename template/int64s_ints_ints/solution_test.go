package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums    []int
		queries []int
		Want    []int64
	}

	testGroup := []testCase{
		{[]int{3, 1, 6, 8}, []int{1, 5}, []int64{14, 10}},
	}

	for i, v := range testGroup {
		result := minOperations(common.DeepCopy(v.nums), common.DeepCopy(v.queries))
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
