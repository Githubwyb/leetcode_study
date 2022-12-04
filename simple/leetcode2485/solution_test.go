package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		Want int
	}

	testGroup := []testCase{
		{8, 6},
		{1, 1},
		{4, -1},
	}

	for i, v := range testGroup {
		result := pivotInteger(v.n)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := pivotInteger1(v.n)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := pivotInteger2(v.n)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
