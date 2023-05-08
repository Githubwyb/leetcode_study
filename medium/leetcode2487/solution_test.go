package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		list []int
		Want []int
	}

	testGroup := []testCase{
		{[]int{5, 2, 13, 3, 8}, []int{13, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for i, v := range testGroup {
		tmp := removeNodes(MakeList(v.list))
		result := List2slice(tmp)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		tmp := removeNodes1(MakeList(v.list))
		result := List2slice(tmp)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
