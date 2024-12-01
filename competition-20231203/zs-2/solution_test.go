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
		{[]int{1, 4, 10}, 19, 2},
		{[]int{1, 4, 10, 5, 7, 19}, 19, 1},
		{[]int{1, 1, 1}, 20, 3},
	}

	for i, v := range testGroup {
		result := minimumAddedCoins(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
