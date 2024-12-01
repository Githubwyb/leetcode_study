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

	testGroup := []testCase{
		{"abcabc", [][]int{{1, 1, 3, 5}, {0, 2, 5, 5}}, []bool{true, true}},
	}

	for i, v := range testGroup {
		result := canMakePalindromeQueries(v.s, DeepCopyIntss(v.queries))
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
