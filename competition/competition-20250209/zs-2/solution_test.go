package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		groups   []int
		elements []int
		Want     []int
	}

	testGroup := []testCase{
		{groups: []int{8, 4, 3, 2, 4}, elements: []int{4, 2}, Want: []int{0, 0, -1, 1, 0}}, // 案例1
		{groups: []int{2, 3, 5, 7}, elements: []int{5, 3, 3}, Want: []int{-1, 1, 0, -1}},   // 案例2
		{groups: []int{10, 21, 30, 41}, elements: []int{2, 1}, Want: []int{0, 1, 0, 1}},    // 案例3
	}

	for i, v := range testGroup {
		result := assignElements(DeepCopy(v.groups), DeepCopy(v.elements))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
