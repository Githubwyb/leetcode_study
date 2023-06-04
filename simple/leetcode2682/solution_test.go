package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		k    int
		Want []int
	}

	testGroup := []testCase{
		{5, 2, []int{4, 5}},
		{4, 4, []int{2, 3, 4}},
	}

	for i, v := range testGroup {
		result := circularGameLosers(v.n, v.k)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
