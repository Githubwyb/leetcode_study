package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		input string
		Want  string
	}

	testGroup := []testCase{
		{"c@da eaf", "a@ac def"},
	}

	for i, v := range testGroup {
		result := sortNormalChar(v.input)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
