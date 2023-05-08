package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		reward1 []int
		reward2 []int
		k       int
		Want    int
	}

	testGroup := []testCase{
		{[]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2, 15},
		{[]int{1, 1}, []int{1, 1}, 2, 2},
	}

	for i, v := range testGroup {
		result := miceAndCheese(common.DeepCopy(v.reward1), common.DeepCopy(v.reward2), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
