package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s        string
		distance []int
		Want     bool
	}

	testGroup := []testCase{
		{"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
		{"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}

	for i, v := range testGroup {
		result := checkDistances(v.s, v.distance)
		if result != v.Want {
			t.Fatalf("%d, v.s %v v.distance %v expect '%v' but '%v'", i, v.s, v.distance, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
