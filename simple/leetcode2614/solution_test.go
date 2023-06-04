package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		Want int
	}

	testGroup := []testCase{
		{[][]int{{1,2,3},{5,6,7},{9,10,11}}, 11},
		{[][]int{{1,2,3},{5,6,7},{9,10,11}}, 11},
		{[][]int{{1,2,3},{5,17,7},{9,11,10}}, 17},
	}

	for i, v := range testGroup {
		result := diagonalPrime(common.DeepCopyIntss(v.grid))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
