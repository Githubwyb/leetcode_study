package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		input string
		Want  []string
	}

	testGroup := []testCase{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for i, v := range testGroup {
		result := getAllString(v.input)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		}
		for i := range result {
			if result[i] != v.Want[i] {
				t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
}
