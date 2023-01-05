package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		left  int
		right int
		Want  []int
	}

	testGroup := []testCase{
		{999998, 1000000, []int{-1, -1}},
		{84084, 407043, []int{84179, 84181}},
		{1, 1000000, []int{2, 3}},
		{1, 1, []int{-1, -1}},
		{10, 19, []int{11, 13}},
		{4, 6, []int{-1, -1}},
	}

	for i, v := range testGroup {
		result := closestPrimes(v.left, v.right)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := closestPrimes1(v.left, v.right)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
