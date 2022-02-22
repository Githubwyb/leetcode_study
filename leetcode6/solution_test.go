package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s      string
		numRow int
		Want   string
	}

	testGroup := []*testCase{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for i, v := range testGroup {
		result := convert(v.s, v.numRow)
		if result != v.Want {
			t.Fatalf("%d, s '%v' numRow %v, expect '%v' but '%v'", i, v.s, v.numRow, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
