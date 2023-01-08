package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want bool
	}

	testGroup := []testCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for i, v := range testGroup {
		result := isValid(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
