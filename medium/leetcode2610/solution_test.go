package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want [][]int
	}

	testGroup := []testCase{
		{[]int{1, 3, 4, 1, 2, 3, 1}, [][]int{{1, 3, 4, 2}, {1, 3}, {1}}},
		{[]int{1, 2, 3, 4}, [][]int{{4, 3, 2, 1}}},
	}

	for i, v := range testGroup {
		result := findMatrix(common.DeepCopy(v.nums))
		if !checkArray(result, v.nums) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}

func checkArray(arr [][]int, nums []int) bool {
	checkMap := make(map[int]bool)
	for _, v := range nums {
		checkMap[v] = true
	}
	for _, v := range arr {
		checkMap1 := make(map[int]bool)
		for _, v1 := range v {
			if !checkMap[v1] {
				return false
			}
			if checkMap1[v1] {
				return false
			}
			checkMap1[v1] = true
		}
	}
	return true
}
