package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		i    string
		p1   [][]string
		r1   []float64
		p2   [][]string
		r2   []float64
		Want float64
	}

	testGroup := []testCase{}

	for i, v := range testGroup {
		result := maxAmount(v.i, DeepCopySlices(v.p1), DeepCopy(v.r1), DeepCopySlices(v.p2), DeepCopy(v.r2))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
