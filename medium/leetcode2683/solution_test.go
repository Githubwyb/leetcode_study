package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want bool
	}

	testGroup := []testCase{
		{[]int{1, 1, 0}, true},
		{[]int{1, 1}, true},
		{[]int{1, 0}, false},
	}

	for i, v := range testGroup {
		result := doesValidArrayExist(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
