package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		x    int
		Want int
	}

	testGroup := []*testCase{
		{123, 321},
		{1534236469, 0},
	}

	for i, v := range testGroup {
		result := reverse(v.x)
		if result != v.Want {
			t.Fatalf("%d, x %v expect '%v' but '%v'", i, v.x, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
