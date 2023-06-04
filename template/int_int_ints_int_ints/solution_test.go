package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n      int
		p      int
		banned []int
		k      int
		Want   []int
	}

	testGroup := []testCase{
		{4, 0, []int{1, 2}, 4, []int{0, -1, -1, 1}},
	}

	for i, v := range testGroup {
		result := minReverseOperations(v.n, v.p, DeepCopy(v.banned), v.k)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
