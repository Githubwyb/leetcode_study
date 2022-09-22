package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		m    int
		Want bool
	}

	testGroup := []testCase{
		{0, 1, true},
		{0, 2, true},
		{0, 100, true},
		{1, 1, false},
		{1, 2, true},
		{1, 100, true},
		{2, 2, false},
		{2, 3, true},
		{2, 100, true},
		{3, 3, false},
		{3, 4, true},
		{3, 100, true},
	}

	for i, v := range testGroup {
		result := checkIsFirstWin(v.n, v.m)
		if result != v.Want {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
