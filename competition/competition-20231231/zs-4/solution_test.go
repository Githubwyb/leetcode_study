package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s       string
		queries [][]int
		Want    []bool
	}

	/*
			"aacaeaacea"
		[[0,0,8,8},{1,4,8,9},{0,3,5,5},{2,4,6,8},{0,1,6,9},{2,4,8,8]]
	*/
	testGroup := []testCase{
		{"aacaeaacea", [][]int{{0, 0, 8, 8}, {1, 4, 8, 9}, {0, 3, 5, 5}, {2, 4, 6, 8}, {0, 1, 6, 9}, {2, 4, 8, 8}}, []bool{false, true, false, true, false, false}},
		{"abcabc", [][]int{{1, 1, 3, 5}, {0, 2, 5, 5}}, []bool{true, true}},
		{"bbaabb", [][]int{{0, 1, 4, 5}, {0, 2, 3, 5}, {2, 2, 5, 5}}, []bool{true, true, true}},
		{"abbcdecbba", [][]int{{0, 2, 7, 9}}, []bool{false}},
		{"acbcab", [][]int{{1, 2, 4, 5}}, []bool{true}},
	}

	for i, v := range testGroup {
		result := canMakePalindromeQueries(v.s, DeepCopyIntss(v.queries))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
