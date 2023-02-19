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
		{39, 3},
		{54, 3},
	}

	for i, v := range testGroup {
		result := minOperations(v.n)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
