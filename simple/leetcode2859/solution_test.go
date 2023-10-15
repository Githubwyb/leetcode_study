package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int
	}

	testGroup := []testCase{
		{[]int{5,10,1,5,2}, 1, 13},
		{[]int{4,3,2,1}, 2, 1},
	}

	for i, v := range testGroup {
		result := sumIndicesWithKSetBits(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
