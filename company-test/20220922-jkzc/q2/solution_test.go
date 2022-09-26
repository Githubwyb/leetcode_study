package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		m    int
		k    int
		Want bool
	}

	testGroup := []testCase{
		{0, 1, 1, true},
		{0, 2, 1, false},
		{0, 2, 2, true},
		{0, 3, 2, false},
		{1, 3, 1, false},
		{1, 3, 4, true},
	}

	for i, v := range testGroup {
		result := checkIsFirstWin(v.n, v.m, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
