package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		num  int
		Want bool
	}

	testGroup := []testCase{
		{443, true},
		{63, false},
		{181, true},
	}

	for i, v := range testGroup {
		result := sumOfNumberAndReverse(v.num)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
