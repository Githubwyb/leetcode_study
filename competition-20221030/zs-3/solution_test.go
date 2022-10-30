package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n      int64
		target int
		Want   int64
	}

	testGroup := []testCase{
		{16, 6, 4},
		{467, 6, 33},
		{1, 1, 0},
		{14674, 6, 326},
	}

	for i, v := range testGroup {
		result := makeIntegerBeautiful(v.n, v.target)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
