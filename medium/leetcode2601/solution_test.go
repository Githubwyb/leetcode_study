package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want bool
	}

	testGroup := []testCase{
		{[]int{4, 9, 6, 10}, true},
		{[]int{6, 8, 11, 12}, true},
		{[]int{5, 8, 3}, false},
	}

	for i, v := range testGroup {
		result := primeSubOperation(common.DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := primeSubOperation1(common.DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
