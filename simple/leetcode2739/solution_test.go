package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n    int
		time int
		Want int
	}

	testGroup := []testCase{
		{5, 10, 60},
		{1, 2, 10},
	}

	for i, v := range testGroup {
		result := distanceTraveled(v.n, v.time)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
