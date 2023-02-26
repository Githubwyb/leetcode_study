package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		word string
		m    int
		Want []int
	}

	testGroup := []testCase{
		{"91221181269244172125025075166510211202115152121212341281327", 21, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{"998244353", 3, []int{1, 1, 0, 0, 0, 1, 1, 0, 0}},
		{"1010", 10, []int{0, 1, 0, 1}},
	}

	for i, v := range testGroup {
		result := divisibilityArray(v.word, v.m)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
