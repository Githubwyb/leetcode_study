package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		k    int
		Want int
	}

	testGroup := []testCase{
		{"165462", 60, 4},
		{"238182", 5, -1},
	}

	for i, v := range testGroup {
		result := minimumPartition(v.s, v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
