package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		Want []int
	}

	testGroup := []testCase{
		{17, []int{2, 0}},
		{2, []int{0, 1}},
	}

	for i, v := range testGroup {
		result := evenOddBit(v.n)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
