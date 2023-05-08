package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		numOnes    int
		numZeros   int
		numNegOnes int
		k          int
		Want       int
	}

	testGroup := []testCase{
		{3, 2, 0, 2, 2},
	}

	for i, v := range testGroup {
		result := kItemsWithMaximumSum(v.numOnes, v.numZeros, v.numNegOnes, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
