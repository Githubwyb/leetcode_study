package main

import (
	"fmt"
	"testing"
	. "leetcode/common"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		price []int
		k     int
		Want  int
	}

	testGroup := []testCase{
		{[]int{13, 5, 1, 8, 21, 2}, 3, 8},
		{[]int{1, 3, 1}, 2, 2},
		{[]int{7, 7, 7, 7}, 2, 0},
	}

	for i, v := range testGroup {
		result := maximumTastiness(DeepCopy(v.price), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
