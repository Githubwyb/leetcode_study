package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}, Want: 2}, // 第一个case
		{nums: []int{4, 5, 6, 4, 4}, Want: 2},             // 第二个case
		{nums: []int{6, 7, 8, 9}, Want: 0},                // 第三个case
	}

	for i, v := range testGroup {
		result := minimumOperations(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
