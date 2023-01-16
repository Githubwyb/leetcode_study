package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int64
	}

	testGroup := []testCase{
		{[]int{10, 10, 10, 10, 10}, 5, 50},
		{[]int{1, 10, 3, 3, 3}, 3, 17},
	}

	for i, v := range testGroup {
		result := maxKelements(common.DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
