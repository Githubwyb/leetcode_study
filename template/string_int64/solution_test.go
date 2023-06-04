package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want int64
	}

	testGroup := []testCase{
		{"0011", 2},
	}

	for i, v := range testGroup {
		result := minimumCost(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
