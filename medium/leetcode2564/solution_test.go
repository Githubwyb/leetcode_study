package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s       string
		queries [][]int
		Want    [][]int
	}

	testGroup := []testCase{
		{"101101", [][]int{{0, 5}, {1, 2}}, [][]int{{0, 2}, {2, 3}}},
		{"0101", [][]int{{12, 8}}, [][]int{{-1, -1}}},
		{"1", [][]int{{4, 5}}, [][]int{{0, 0}}},
	}

	for i, v := range testGroup {
		result := substringXorQueries(v.s, v.queries)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := substringXorQueries1(v.s, v.queries)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
