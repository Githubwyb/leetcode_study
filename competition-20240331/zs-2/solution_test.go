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
		{13, 6, 15},
		{10, 3, 13},
	}

	for i, v := range testGroup {
		result := maxBottlesDrunk(v.n, v.time)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
