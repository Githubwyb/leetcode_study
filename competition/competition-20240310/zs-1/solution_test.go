package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums     []int
		divisors []int
		Want     int
	}

	testGroup := []testCase{
		{[]int{1, 3, 2}, []int{4, 3, 1, 5, 2}, 2},
		{[]int{5, 5, 5}, []int{2, 4, 2, 7}, 4},
	}

	for i, v := range testGroup {
		result := minimumBoxes(DeepCopy(v.nums), DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
