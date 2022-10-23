package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int32
		k    int32
		Want int32
	}

	testGroup := []testCase{
		{13, 2, 10},
		{1, 1, 1},
		{math.MaxInt32, math.MaxInt32, 1398283264},
	}

	for i, v := range testGroup {
		result := findKthNumber(v.n, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
