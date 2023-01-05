package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums int
		Want int
	}

	testGroup := []testCase{
		{7, 1},
		{121, 2},
		{1248, 4},
	}

	for i, v := range testGroup {
		result := countDigits(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
