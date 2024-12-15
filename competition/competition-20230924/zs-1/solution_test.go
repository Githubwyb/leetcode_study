package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want string
	}

	testGroup := []testCase{
		{"010", "001"},
		{"0101", "1001"},
	}

	for i, v := range testGroup {
		result := maximumOddBinaryNumber(v.s)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
